build:
	go build -ldflags "-s -w" -o ./emi ./cmd/emi && \
	make build-js-sdks && \
	make build-envelopes && \
	go build -ldflags "-s -w" -o ./emi ./cmd/emi && \
	make wasm && \
	./emi spec --output .vscode/ && \
	./emi spec --output ./playground/public/

win:
	go build -ldflags "-s -w" -o ./emi.exe ./cmd/emi

unix:
	go build -ldflags "-s -w" -o ./emi ./cmd/emi

build-js-sdks:
	cd js-sdk-kit && npm run build && rm -rf ../lib/js/ts-sdk && cp -R build ../lib/js/ts-sdk && \
	cd - && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	rm -rf ./lib/js/ts-sdk && cp -R ./js-sdk-kit/src ./lib/js/ts-sdk && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go

build-envelopes:
	cd envelopes && npm run compile && cd -  && \
	rm -rf ./lib/js/ts-envelopes && \
	cp -R ./envelopes/src ./lib/js/ts-envelopes && \
	cp ./lib/js/index.go.txt ./lib/js/ts-envelopes/index.go && \
	rm -rf ./lib/js/ts-envelopes && cp -R ./envelopes/src ./lib/js/ts-envelopes && \
	cp ./lib/js/index.go.txt ./lib/js/ts-envelopes/index.go

all: 
	make build && make build-envelopes && make jstests && make sample
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

