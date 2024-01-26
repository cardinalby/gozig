# gozig

The tool reads `GOOS` and `GOARCH` env vars and passes the corresponding `target` arguments 
to `zig cc` and `zig c++` invocations.

To make it work with `go build` you need to set the following env vars:
- `CC=gozig cc`
- `CXX=gozig c++`

Original idea and implementation is from [zigtool](https://github.com/dosgo/zigtool)

This tool is used in [go-build-everywhere](https://github.com/cardinalby/go-build-everywhere) toolset.

# Targets support
`linux`, `windows`, `darwin` and `wasm` targets are supported.

# Platform-specific flags
To add MacOS SDK path use `GOZIG_OSX_SDK` env variable (will be used for `darwin` builds).

## Installation

```bash
go install github.com/cardinalby/gozig@latest
```

### Debug invocations
Set `DEBUG_EXEC=1` env variable to see the actual `zig cc` and `zig c++` invocation arguments in the output.