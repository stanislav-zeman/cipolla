package initor

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/stanislav-zeman/cipolla/internal/config"
)

const (
	defaultFilePermissions = 0o755
	defaultClientTimeout   = 5 * time.Second
)

var errUnexpectedStatusCode = errors.New("unexpected status code when getting data")

// Initor inits projects metadata and configuration such as
// the Go module itself, gitignore, golangci configuration etc.
type Initor struct {
	config    config.Config
	directory string
	client    http.Client
}

func New(config config.Config, directory string) Initor {
	client := http.Client{
		Timeout: defaultClientTimeout,
	}

	return Initor{
		config:    config,
		directory: directory,
		client:    client,
	}
}

func (i *Initor) Run(ctx context.Context) error {
	err := i.initGoModule(ctx)
	if err != nil {
		return err
	}

	if i.config.Misc.Gitignore != "" {
		err := i.initGitignore(ctx)
		if err != nil {
			return fmt.Errorf("failed initializing gitignore: %w", err)
		}
	}

	if i.config.Misc.Makefile != "" {
		err := i.initMakefile(ctx)
		if err != nil {
			return fmt.Errorf("failed initializing makefile: %w", err)
		}
	}

	if i.config.Misc.GolangCI != "" {
		err := i.initGolangCI(ctx)
		if err != nil {
			return fmt.Errorf("failed initializing golangci: %w", err)
		}
	}

	return nil
}

func (i *Initor) AddDependencies(ctx context.Context, dependencies []string) error {
	for _, dependency := range dependencies {
		if dependency == "" {
			continue
		}

		// No need to import standard library packages.
		if !strings.Contains(dependency, ".") {
			continue
		}

		cmd := exec.CommandContext(ctx, "go", "get", dependency)
		cmd.Dir = i.directory

		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed adding dependency to go module (%s): %w", out, err)
		}
	}

	cmd := exec.CommandContext(ctx, "go", "mod", "tidy")
	cmd.Dir = i.directory

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to tidy go module (%s): %w", out, err)
	}

	return nil
}

func (i *Initor) initGoModule(ctx context.Context) error {
	path := filepath.Join(i.directory, "go.mod")

	// Check if go.mod already exists.
	// The go mod init command fails if that is the case.
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to stat go module: %w", err)
	}

	err = os.MkdirAll(i.directory, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("failed creating project directory: %w", err)
	}

	cmd := exec.CommandContext(ctx, "go", "mod", "init", i.config.Module)
	cmd.Dir = i.directory

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed initializing go module (%s): %w", out, err)
	}

	return nil
}

func (i *Initor) initGitignore(ctx context.Context) error {
	path := filepath.Join(i.directory, ".gitignore")

	err := i.httpWriteFile(ctx, i.config.Misc.Gitignore, path)
	if err != nil {
		return fmt.Errorf("failed to http write file: %w", err)
	}

	return nil
}

func (i *Initor) initGolangCI(ctx context.Context) error {
	path := filepath.Join(i.directory, ".golangci.yaml")

	err := i.httpWriteFile(ctx, i.config.Misc.GolangCI, path)
	if err != nil {
		return fmt.Errorf("failed to http write file: %w", err)
	}

	return nil
}

func (i *Initor) initMakefile(ctx context.Context) error {
	path := filepath.Join(i.directory, "Makefile")

	err := i.httpWriteFile(ctx, i.config.Misc.Makefile, path)
	if err != nil {
		return fmt.Errorf("failed to http write file: %w", err)
	}

	return nil
}

func (i *Initor) httpWriteFile(ctx context.Context, url, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed creating request: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := i.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed getting data: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errUnexpectedStatusCode, resp.StatusCode)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed reading reader: %w", err)
	}

	err = os.WriteFile(path, data, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("failed writing file: %w", err)
	}

	return nil
}
