#!/bin/bash

npm run build
live-server ../../dist --ignore=go &

fswatch -o -e ".*" -i "\\.templ$" . | while read; \
    do \
        echo "🚀 Updating templ"
        templ generate
        GOARCH=wasm GOOS=js go build -o ../public/main.wasm main.go
        npm run build
        #cd ../../
        #tailwindcss build -m -o src/css/style.css
        #cd src/go
    done
