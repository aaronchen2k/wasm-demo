default: cp build run

run:
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

cp:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static

build: main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/main.wasm .