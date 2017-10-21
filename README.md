# gosif

String Into File: Insert strings into a file in a given line.


[![Build Status](https://travis-ci.org/chentex/go-sif.svg)](https://travis-ci.org/chentex/go-sif)
[![codecov](https://codecov.io/gh/chentex/go-sif/branch/master/graph/badge.svg)](https://codecov.io/gh/chentex/go-sif/branch/master)
[![GoDoc](https://godoc.org/github.com/chentex/go-sif?status.svg)](https://godoc.org/github.com/chentex/go-sif)
[![Go Report Card](https://goreportcard.com/badge/github.com/chentex/go-sif)](https://goreportcard.com/report/github.com/chentex/go-sif)


## Installation

Find binaries for linux and MacOS in the [releases page](https://github.com/chentex/go-sif/releases)

## Getting Started

### Insert line in a file

```bash
gosif insert -f /your/path/to/file.txt -l2 -t "The text you want to insert"
```

### Default values

```bash
-l, --line = default value will insert at the end of the file
-t, --text = default value will insert an empty line
```

## Commands Usage

### Insert

```bash
Usage:
  gosif insert [flags]

Global Flags:
  -f, --file string   The file where you want to insert the string
  -l, --line int      Line where the string is going to be inserted, if omited the string will be inserted
 in a new line at the end of the file. Count starts at 1. (default -1)
  -t, --text string   The to insert into file
```

### Version

```bash
Usage:
  gosif version
```
