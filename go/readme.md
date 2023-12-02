# wasm

```sh
GOARCH=wasm GOOS=js go build -o lib.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../poc/src/js/
mkdir -p ../poc/src/wasm/
cp lib.wasm ../poc/src/wasm/
```