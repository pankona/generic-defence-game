
wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

upload:
	gcloud storage cp main.wasm gs://generic-defence-game/main.wasm
