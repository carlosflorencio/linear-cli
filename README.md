# Linear CLI

The Linear CLI is a command-line interface for interacting with the [Linear](https://linear.app/) project management tool.

## Requirements

- go
- Linear API Key

## Install

```shell
go install github.com/carlosflorencio/linear-cli
```

## Set the API Key

```shell
# Required shell environment variable
export LINEAR_API_KEY=lin_api_8U61z4PtRUJ24gd394c2BAJF
```

## Usage

```shell
Usage:
  linear [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  user        Prints information about the user

Flags:
  -h, --help          help for linear
  -t, --team string   Team ID

Use "linear [command] --help" for more information about a command.
```

## Status

Work in Progress.
