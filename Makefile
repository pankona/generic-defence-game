
wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

upload:
	gcloud storage cp main.wasm gs://generic-defence-game/main.wasm
	gcloud storage objects update gs://generic-defence-game/main.wasm --add-acl-grant=entity=AllUsers,role=READER
