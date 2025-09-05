build:
	cd js-sdk-kit && npm run build && rm -rf ../lib/js/ts-sdk && cp -R build ../lib/js/ts-sdk;
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	rm -rf ./lib/js/ts-sdk && cp -R ./js-sdk-kit/src ./lib/js/ts-sdk && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	go build -ldflags "-s -w" -o ./emi ./cmd/emi && \
	make wasm && \
	./emi spec --output .vscode/emi-definitions.json && \
	./emi spec --output ./playground/public/emi-definitions.json

all: 
	make build && make jstests && make sample
sample:
	cd examples/js-test/reactclient && make && cd -

wasm:
	GOOS=js GOARCH=wasm go build -o ./playground/public/emi-compiler.wasm ./cmd/emi-wasm/main.go && \
	cp ./playground/public/emi-compiler.wasm  ./emi-npm/bin/emi-compiler.wasm

jstests:
	cd tests/js && npm i && npx vitest run
ci:
	make build;
	cd playground && npm run build && cd -;
	make sample;
	make jstests;

compile-github:
	rm -rf __webdir && cp -R emi-web/build __webdir && touch __webdir/.nojekyll && cp -R playground/dist __webdir/playground