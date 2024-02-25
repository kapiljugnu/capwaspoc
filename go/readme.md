# wasm

```sh
GOARCH=wasm GOOS=js go build -o main.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../cap/src/public/
cp main.wasm ../cap/src/public/
```