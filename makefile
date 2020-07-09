default: build

test:
	go test maps

build:
	GOARCH=wasm GOOS=js go build -o lib.wasm main

color-correction:
	go build -o color-correction.exe correction
