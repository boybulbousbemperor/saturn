You can compile the Rust source code into WebAssembly bytecode using the command below.
```SH build-wasm.sh
GOARCH=wasm GOOS=js go build -o web/app.wasm
```