package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stanislav-zeman/cipolla/internal/config"
	"github.com/stanislav-zeman/cipolla/internal/initor"
	processor "github.com/stanislav-zeman/cipolla/internal/procesor"
	"github.com/stanislav-zeman/cipolla/internal/templator"
	"github.com/stanislav-zeman/cipolla/internal/writer"
	yaml "gopkg.in/yaml.v3"
)

const applicationTimeout = 30 * time.Second

var (
	configPath         = flag.String("config", "cipolla.yaml", "project structure configuration")
	outputDirectory    = flag.String("out", ".", "project structure output directory")
	templatesDirectory = flag.String("templates", "assets", "directory with templates")
)

func main() {
	flag.Parse()

	log.Println("running cipolla...")

	err := runCipolla()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func runCipolla() error {
	f, err := os.ReadFile(*configPath)
	if err != nil {
		return fmt.Errorf("failed reading config file: %w", err)
	}

	var conf config.Config

	err = yaml.Unmarshal(f, &conf) //nolint: musttag
	if err != nil {
		return fmt.Errorf("failed unmarshalling config file: %w", err)
	}

	i := initor.New(conf, *outputDirectory)

	ctx, cancel := context.WithTimeout(context.Background(), applicationTimeout)
	defer cancel()

	err = i.Run(ctx)
	if err != nil {
		return fmt.Errorf("failed running initor: %w", err)
	}

	t, err := templator.New(*templatesDirectory)
	if err != nil {
		return fmt.Errorf("failed creating templator: %w", err)
	}

	w := writer.NewWriter(*outputDirectory)
	p := processor.New(conf, t, w)

	dependencies, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed running processor: %w", err)
	}

	err = i.AddDependencies(ctx, dependencies)
	if err != nil {
		return fmt.Errorf("failed adding dependencies to go module via initor: %w", err)
	}

	return nil
}
