# fsserver

[![Build Status](https://github.com/goasm/fsserver/actions/workflows/go.yml/badge.svg)](https://github.com/goasm/fsserver/actions/workflows/go.yml)

fsserver is a command-line static HTTP server.

## Installation

```
go install github.com/goasm/fsserver/cmd/fsserver@latest
```

## Usage

```
fsserver [OPTION...] PATH
```

| Option    | Description            |
| --------- | ---------------------- |
| -h/--help | print help message     |
| -a        | address to use         |
| -p        | port to bind to        |
