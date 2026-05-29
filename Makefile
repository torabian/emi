build:
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
	cd ./examples/js-sdk-kit && npm run build && rm -rf ../../lib/js/ts-sdk && cp -R build ../../lib/js/ts-sdk && \
	cd - && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	rm -rf ./lib/js/ts-sdk && cp -R ./examples/js-sdk-kit/src ./lib/js/ts-sdk && \
	rm -rf ./lib/js/js-sdk && cp -R ./examples/js-sdk-kit/build ./lib/js/js-sdk && \
	cp ./lib/js/index.go.txt ./lib/js/ts-sdk/index.go && \
	cp ./lib/js/index.go.txt ./lib/js/js-sdk/index.go

build-envelopes:
	cd ./examples/envelopes && npm run compile && npm run build && cd -  && \
	rm -rf ./lib/js/ts-envelopes && \
	rm -rf ./lib/js/js-envelopes && \
	cp -R ./examples/envelopes/src ./lib/js/ts-envelopes && \
	cp -R ./examples/envelopes/build ./lib/js/js-envelopes && \
	cp ./lib/js/index.go.txt ./lib/js/ts-envelopes/index.go && \
	rm -rf ./lib/js/ts-envelopes && cp -R ./examples/envelopes/src ./lib/js/ts-envelopes && \
	cp ./lib/js/index.go.txt ./lib/js/ts-envelopes/index.go && \
	cp ./lib/js/index.go.txt ./lib/js/js-envelopes/index.go

all: 
	make build && make build-envelopes && make jstests && make sample && make nullabletest
sample:
	cd examples/js-test/reactclient && make && cd -

wasm:
	GOOS=js GOARCH=wasm go build -o ./playground/public/emi-compiler.wasm ./cmd/emi-wasm/main.go && \
	cp ./playground/public/emi-compiler.wasm  ./examples/emi-wasm-helper/emi-compiler.wasm

jstests:
	cd examples/js && npm i && npx vitest run
ci:
	make build;
	cd playground && npm run build && cd -;
	make sample;
	make jstests;

compile-github:
	rm -rf __webdir && \
	cp -R ./examples/emi-web/build __webdir && touch __webdir/.nojekyll && \
	cp -R playground/dist __webdir/playground && \
	cp -Rexamples/in-browser-server/browser __webdir/in-browser-server


qpsamples:
	./emi qp dir --path examples/query-predict/queries-source --output examples/query-predict/output

win:
	go build -ldflags "-s -w" -o ./emi.exe ./cmd/emi
go:

	./emi.exe go --path ./examples/fullstack/definitions.emi.yml --output ./examples/fullstack/sdk --emigo github.com/torabian/emi/examples/fullstack/emigo

nullabletest:
	go test ./examples/nullable-test/...

collectiontest:
	go test -v ./tests/collection/... && go test -v ./tests/array/...


test_examples:
	cd examples/fullstack && make