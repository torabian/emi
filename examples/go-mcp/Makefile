def:
	../../emi go --path ./gomcp.emi.yml --output ./internal/defs --pkg defs

server:
	go build -o ./gomcp-server ./cmd/server

run: server
	./gomcp-server
