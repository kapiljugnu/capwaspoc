# capacitor-wasm proof of concept

## live reload of templ

```sh
cd src/go
./watch.sh
```

## building and testing

### go stuff

```sh
# compile go wasm stuff
cd go
GOARCH=wasm GOOS=js go build -o main.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../src/public/
cp main.wasm ../src/public/
```

### watch go stuff
```sh
cd src/go
reflex -c reflex.conf
```

### capacitor stuff

```sh
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
npx cap add ios
npx cap run ios
```

## deploy to ionic applow

```sh
npx cap add android
git push ionic
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

## ionic appflow stuff

```sh
ionic init
ionic link 17c1cf69
```
