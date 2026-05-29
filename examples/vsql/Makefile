EMI ?= ../../emi

generate:
	$(EMI) go --path ./createuser.emi.yml --output ./dto --pkg dto && \
	$(EMI) js --path ./createuser.emi.yml --output ./js

run: generate
	go run .

test:
	go test -v ./tests/...

.PHONY: generate run test
