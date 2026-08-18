compile:
	./gradlew build

test:
	./gradlew test

gen: gen-sdkgen gen-inventory gen-billing

gen-sdkgen:
	../../emi kotlin --path ./definitions.emi.yml --output ./app/src/main/kotlin/org/example/sdkgen

# --pkg must match the package the other module's `module:` fields reference it by -
# see inventory.emi.yml/billing.emi.yml. --tags android-forms also emits a
# <Dto>FormState alongside every inventory dto (see lib/kotlin/kotlin-form-state.go).
# no-sdk on both: gen-sdkgen's own output already carries the shared emikot runtime
# (MaybeField/Maybe/GResponse/EmiWebSocketX/...) once for the whole compilation unit -
# see --tags no-sdk in `emi tags`.
gen-inventory:
	../../emi kotlin --path ./inventory.emi.yml --output ./app/src/main/kotlin/org/example/inventory --pkg org.example.inventory --tags android-forms,no-sdk

gen-billing:
	../../emi kotlin --path ./billing.emi.yml --output ./app/src/main/kotlin/org/example/billing --pkg org.example.billing --tags no-sdk

clean:
	rm -rf ./app/src/main/kotlin/org/example/sdkgen ./app/src/main/kotlin/org/example/inventory ./app/src/main/kotlin/org/example/billing

build:
	make gen && make compile
