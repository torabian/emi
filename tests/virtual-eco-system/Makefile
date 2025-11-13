genall:
	make genkotlin && make genjs && make gengo && make gents

genkotlin:
	../../emi kotlin --path ./virtual.emi.yml --output ./kotlin/app/src/main/kotlin/org/example/sdkgen

genjs:
	../../emi js --path ./virtual.emi.yml --output ./js/sdkgen

gents:
	../../emi js --path ./virtual.emi.yml --output ./ts/sdkgen --tags typescript

clean: 
	rm -rf ./js/sdkgen && rm -rf ./ts/sdkgen && rm -rf ./kotlin/app/src/main/kotlin/org/example/sdkgen && rm -rf ./go/sdkgen

gengo:
	../../emi go --path ./virtual.emi.yml --output ./go/sdkgen  --emigo test/emi/emigo

buildsdk:
	cd kotlin && make


build:
	make clean && \
	make genall && \
	cd kotlin && make