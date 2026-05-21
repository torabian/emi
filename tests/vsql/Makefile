EMI ?= ../emi

generate:
	$(EMI) go --path ./createuser.emi.yml --output ./dto --pkg dto

run: generate
	go run .

test:
	go test -v ./tests/...

.PHONY: generate run test
