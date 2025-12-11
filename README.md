[![Go Report Card](https://goreportcard.com/badge/github.com/stanislav-zeman/cipolla)](https://goreportcard.com/report/github.com/stanislav-zeman/cipolla)
![Go version](https://img.shields.io/github/go-mod/go-version/stanislav-zeman/cipolla)
![CI status](https://github.com/stanislav-zeman/cipolla/actions/workflows/ci.yaml/badge.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/stanislav-zeman/cipolla/badge.svg?branch=master)](https://coveralls.io/github/stanislav-zeman/cipolla?branch=master)

# Cipolla

The idea of this project is to have a tool that generates project structure
following the Onion/Hexagonal/Ports & Adapters architecture principles.

The tool is configured using YAML config that is then parsed and templated.
The goal is not to create a completely functional project but to generate most
of the boilerplate code and files that you then can go and easily edit.

Apart from generating things like Domain objects, Application services,
API controllers etc., the tool can also bootstrap project files like Makefile,
GolangCI linter configuration, gitignore etc.

## Usage

```sh
cipolla --templates ./cipolla/assets --config ./cipolla/config/example.yaml --out my-project
```

The project provides example configuration and showcases the generated
code in the `example` directory.

To view the example project directory structure checkout the
generated [project structure file](docs/example-structure.txt).
