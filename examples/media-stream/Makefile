# Regenerates both sides from media.emi.yml. Needs `emi` on PATH (or go run ../../cmd/emi).
.PHONY: gen device device-c web run-server
EMI ?= emi

gen:
	$(EMI) go  --tags skip-cli,skip-wasm-gin --path ./media.emi.yml --pkg mediagen --output ./server/lib
	$(EMI) cpp --path ./media.emi.yml --output ./device/emi
	$(EMI) js  --tags typescript --path ./media.emi.yml --output ./web/sdk
	$(EMI) c   --path ./media.emi.yml --output ./device-c/emi
	$(EMI) c   --path ./media.emi.yml --output ./esp-idf-mic/main/emi

# Builds the host-side simulator of the ESP32 firmware (fake mic + fake camera).
device:
	$(CC) -c -O1 device/emi/cjson/cJSON.c -o device/cJSON.o
	$(CXX) -std=c++17 -O1 -Wall -Wextra -I device/emi device/main.cpp device/cJSON.o -o device/media-device

# Plain-C client (no libcurl, no cJSON needed: only the websocket + url helpers).
device-c:
	$(CC) -std=c99 -D_POSIX_C_SOURCE=200809L -DFETCHX_NO_CURL -O1 -Wall -Wextra -I device-c/emi device-c/main.c -lm -o device-c/media-device-c

# Bundles the browser demo into web/public/app.js (needs npm; installs qs + esbuild).
web:
	cd web && npm install && npm run build

run-server:
	cd server && go run .
