build:
	make build-js-sdks && \
	make build-envelopes && \
	go build -ldflags "-s -w" -o ./emi ./cmd/emi && \
	make wasm && \
	./emi spec --output .vscode/ && \
	./emi spec --output ./playground/public/


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


release:
	rm -rf ./artifacts/
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ./artifacts/emi cmd/emi/main.go
	cd ./artifacts/ && zip emi_amd64_darwin.zip emi && cd -
# 	packagesbuild --project ./macos-installer.pkgproj && mv ./artifacts/emi.pkg ./artifacts/emi_intel_amd64.pkg
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ./artifacts/emi cmd/emi/main.go
	cd ./artifacts/ && zip emi_arm64_darwin.zip emi && cd -
# 	packagesbuild --project ./macos-installer.pkgproj && mv ./artifacts/emi.pkg ./artifacts/emi_silicon_arm64.pkg
	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ./artifacts/emi.exe cmd/emi/main.go
	cd ./artifacts/ && zip emi_arm64_windows.zip emi.exe && cd -
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ./artifacts/emi.exe cmd/emi/main.go
	cd ./artifacts/ && zip emi_amd64_windows.zip emi.exe && cd -
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ./artifacts/emi cmd/emi/main.go
	cd ./artifacts/ && zip emi_arm64_linux.zip emi && cd -
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ./artifacts/emi cmd/emi/main.go
	cd ./artifacts/ && zip emi_amd64_linux.zip emi && cd -
	rm -rf ./artifacts/emi ./artifacts/emi.exe
	zip -r ./artifacts/emi-node-wasm-package.zip ./emi-npm