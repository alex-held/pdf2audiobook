# Template Repository

![Changelog CI Status](https://github.com/alex-held/template/workflows/auto-changelog/badge.svg)
![GitHub](https://img.shields.io/github/license/alex-held/template?style=flat-square)


# Prerequisites & Tools

To gain maximum benefit of this start template you must have following tools installed locally and in your $PATH.

## General
- git


## Go
- [golang-sdk](https://golang.org/dl/)
- [go-task](https://taskfile.dev/#/installation)
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)


# Quickstart

## Edit Config Files

``` yaml 
# .golangci.yaml

   goimports:
         local-prefixes: <project-import>
```

``` yaml
# Taskfile.yml
vars:
  PROJECT: <project-name> # name of the binary output
```

