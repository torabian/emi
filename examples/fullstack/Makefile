dev:
	../../emi go --path ./definitions.emi.yml --output ./sdk --emigo github.com/torabian/emi/examples/fullstack/emigo && \
	../../emi js --path ./definitions.emi.yml --output ./client
	go build -ldflags "-s -w" -o ./app ./cmd/

refresh:
	cd ../.. && make win && cd - && make dev