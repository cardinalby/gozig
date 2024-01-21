# gozig

The tool reads `GOOS` and `GOARCH` env vars and passes the corresponding `target` arguments 
to `zig cc` and `zig c++` invocations.

To make it work with `go build` you need to set the following env vars:
- `CC=gozig cc`
- `CXX=gozig c++`

Original idea and implementation is from [zigtool](https://github.com/dosgo/zigtool)

# Targets support
`linux`, `windows`, `darwin` and `wasm` targets are supported.

## Installation

```bash
go install github.com/cardinalby/gozig@latest
```