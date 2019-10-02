# Sample project explaining how to use Makefiles to build Golang projects

This is a project which is intended to demonstrate how to use Makefiles to build go projects. Build version tag will be detected automatically.

## Install GNU Make

### Debian based systems

```bash
sudo apt install make
```

### Red Hat based systems

```bash
yum install make
```

## Make Commands configured in this project

This project contains a sample [Makefile](./Makefile) with some build tasks.

```bash
# List configured make tasks
make help

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

## Version tags

- If the source code is exactly on a build tag, the binary will be created with a clean tag. Like `vx.x.x`
- If the source code is modified and have uncommited changes, the build tag would be `vx.x.x-dirty` (latest build tag - dirty)
- If the source code has any untagged changes and the working directory is clean, the build tag will be `vx.x.x-34fdr54`(latest build tag - latest commit SHA)
