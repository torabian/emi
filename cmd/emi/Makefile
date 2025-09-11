
release:
	rm -rf ../../artifacts/
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/emi ./main.go
	cd ../../artifacts/ && zip emi_amd64_darwin.zip emi && cd -
# 	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/emi.pkg ../../artifacts/emi_intel_amd64.pkg
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/emi ./main.go
	cd ../../artifacts/ && zip emi_arm64_darwin.zip emi && cd -
# 	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/emi.pkg ../../artifacts/emi_silicon_arm64.pkg
	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/emi.exe ./main.go
	cd ../../artifacts/ && zip emi_arm64_windows.zip emi.exe && cd -
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/emi.exe ./main.go
	cd ../../artifacts/ && zip emi_amd64_windows.zip emi.exe && cd -
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/emi ./main.go
	cd ../../artifacts/ && zip emi_arm64_linux.zip emi && cd -
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/emi ./main.go
	cd ../../artifacts/ && zip emi_amd64_linux.zip emi && cd -
	rm -rf ../../artifacts/emi ../../artifacts/emi.exe && \
	cd ../.. && zip -r artifacts/emi-node-wasm-package.zip emi-npm