build:
	cd js-sdk-kit && npm run build && rm -rf ../lib/js/prebuilt-sdk && cp -R build ../lib/js/prebuilt-sdk;
	cp ./lib/js/index.go.txt ./lib/js/prebuilt-sdk/index.go && \
	rm -rf ./lib/js/ts-sdk && cp -R ./js-sdk-kit/src ./lib/js/ts-sdk && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	go build -ldflags "-s -w" -o ./emi ./cmd/emi && \
	make wasm


sample:
	make build && \
	./emi js:module --path ./examples/js-test/jsonplaceholder.emi.yml --output ./examples/js-test/backend/src/generated --tags typescript,nestjs

wasm:
	GOOS=js GOARCH=wasm go build -o ./playground/public/emi-compiler.wasm ./cmd/emi-wasm/main.go && \
	cp ./playground/public/emi-compiler.wasm  ./emi-npm/bin/emi-compiler.wasm

jstests:
	cd tests/js && npm i && npm test
ci:
	make build;
	cd playground && npm run build && cd -;
	make sample;
	make jstests;