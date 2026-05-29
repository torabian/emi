def:
	../../emi go --path ./in-browser-server.emi.yml --output ./internal/defs --pkg defs --tags skip-cli,split-gin,no-client,skip-wasm-gin && \
	../../emi js --path ./in-browser-server.emi.yml --output ./browser/gen --tags include-ext

wasm:
	GOOS=js GOARCH=wasm go build -o ./browser/in-browser-server.wasm ./cmd/wasm/main.go 


server:
	go build -o ./browser/server ./cmd/server/main.go 
