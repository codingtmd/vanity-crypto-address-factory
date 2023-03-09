GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/vcaf.wasm .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static