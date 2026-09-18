EMI ?= ../../emi

generate:
	$(EMI) go --path ./queries.emi.yml --output ./sdkgen
	$(EMI) go --path ./predict_join.emi.yml --output ./sdkgen

test: generate
	go test ./...

.PHONY: generate test
