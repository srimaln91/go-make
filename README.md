# Use GNU Make to build Golang projects with embedded build details

This is a project which is intended to demonstrate how to use Makefiles to build go projects. Build version tag will be detected automatically and updated in the compilation time.

[![Build Status](https://travis-ci.org/srimaln91/go-make.svg?branch=master)](https://travis-ci.org/srimaln91/go-make)
[![codecov](https://codecov.io/gh/srimaln91/go-make/branch/master/graph/badge.svg)](https://codecov.io/gh/srimaln91/go-make)
[![Go Report Card](https://goreportcard.com/badge/github.com/srimaln91/go-make)](https://goreportcard.com/report/github.com/srimaln91/go-make)
[![GoDoc](https://godoc.org/github.com/srimaln91/go-make?status.svg)](https://godoc.org/github.com/srimaln91/go-make)

## Install GNU Make

### Debian based systems

```bash
sudo apt install make
```

### Red Hat based systems

```bash
yum install make
```

## How to use

01. Import [Makefile](./Makefile) into your project root
02. Update main function as follows.

```go
package main

import (
    "github.com/srimaln91/go-build"
)

func main() {

    // Print binary details and terminate the program when --version flag provided.
    build.CheckVersion()

    // Starting the bootstrpping process
    // bootstrap.Start()
}

```

## Make Commands configured in this project

This project contains a sample [Makefile](./Makefile) with some build tasks.

```bash
# List configured make tasks
make help

# Run unit tests
make test

# Build project
make build

# Run project
make run

# Clean build directory
make clean
```

## Check Binary Version

```bash
# Can check binary version by passing --version flag
./build/vx.x.x/go-build-linux-amd64 --version
```

```bash
# Output
+----------------+------------+------------------------------------------+-------------+-------------------------+
| BINARY VERSION | GO VERSION |                GIT COMMIT                |   OS/ARCH   |          BUILT          |
+----------------+------------+------------------------------------------+-------------+-------------------------+
| v0.6.0         | go1.12.9   | c8bf7b40e9d842769b580b704931904197e0b713 | linux/amd64 | 2019-10-05-14:01:35-UTC |
+----------------+------------+------------------------------------------+-------------+-------------------------+
```

## Version tags

- If the source code is exactly on a build tag, the binary will be created with a clean tag. Like `vx.x.x`
- If the source code is modified and have uncommited changes, the build tag would be `vx.x.x-dirty` (latest build tag - dirty)
- If the source code has any untagged changes and the working directory is clean, the build tag will be `vx.x.x-34fdr54`(latest build tag - latest commit SHA)
