# fsserver

[![Build Status](https://travis-ci.org/goasm/fsserver.svg?branch=master)](https://travis-ci.org/goasm/fsserver)

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
