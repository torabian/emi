# Regenerates the ESP-IDF server code from thermo.emi.yml. Needs `emi` on PATH (or go run ../../cmd/emi).
.PHONY: gen client
EMI ?= emi

gen:
	$(EMI) c --path ./thermo.emi.yml --output ./esp-idf-thermo/main/emi --tags esp-server

# The matching client (libcurl / websocket) for a host or another device.
client:
	$(EMI) c --path ./thermo.emi.yml --output ./client-c/emi
