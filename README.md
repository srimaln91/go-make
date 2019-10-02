# Sample project explaining how to use Makefiles to build Golang project

This is a project which is intended to demonstrate how to use Makefiles to build go projects. Build version tag will be detected automatically.

## Make Commands

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

- If the source code is exactly on a build tag. The binary will be created with a clean tag. Like `vx.x.x`
- If the source code is modified and have uncommited changes. The build tag will be `vx.x.x-dirty` (latest build tag - dirty)
- If the source code has any untagged changes. The build tag will be `vx.x.x-34fdr54`(latest build tag - latest commit SHA)
