# capacitor-wasm proof of concept

## building and testing

```sh
# compile go wasm stuff
cd go
GOARCH=wasm GOOS=js go build -o main.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../cap/src/public/
cp main.wasm ../cap/src/public/

# capacitor stuff
cd ../cap

# update npm packages
ncu -u
rm package-lock.json node_modules
npm install --verbose
npm audit --omit=dev

# test in browser
npm start

# build for deploy
npm run build

# run in iOS simulator
npx cap sync
npx cap run ios
```
## get iOS local sim working

```sh
rm -rf ios
npm install @capacitor/ios
npx cap add ios
sudo xcode-select -r
```

## devenv stuff

```sh
nix flake init --template github:cachix/devenv#flake-parts
```