EMI ?= ../../emi

generate:
	$(EMI) go --path ./batch_insert_users.emi.yml --output ./sdkgen

test: generate
	go test ./...

generate-ts:
	$(EMI) js --path ./batch_insert_users.emi.yml --output ./ts/sdkgen --tags typescript,no-package

test-ts: generate-ts
	cd ts && npm install --no-audit --no-fund && npm run typecheck

.PHONY: generate test generate-ts test-ts
