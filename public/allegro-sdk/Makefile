gen:
	make gen-kotlin && make gen-js && make gen-ts && make gen-golang

gen-js:
	rm -rf javascript/gen/ && \
	../../emi.exe js --path definitions/offer/user-offer-information.emi.yml --output javascript/gen/offer/user-offer-information && \
	../../emi.exe js --path definitions/offer/offer-translations.emi.yml --output javascript/gen/offer/offer-translations && \
	../../emi.exe js --path definitions/offer/offer-management.emi.yml --output javascript/gen/offer/offer-management

gen-ts:
	rm -rf typescript/gen/ && \
	../../emi.exe js --path definitions/offer/user-offer-information.emi.yml --output typescript/gen/offer/user-offer-information --tags typescript && \
	../../emi.exe js --path definitions/offer/offer-translations.emi.yml --output typescript/gen/offer/offer-translations --tags typescript && \
	../../emi.exe js --path definitions/offer/offer-management.emi.yml --output typescript/gen/offer/offer-management --tags typescript

gen-kotlin:
	rm -rf kotlin/gen/ && \
	../../emi.exe kotlin --path definitions/offer/user-offer-information.emi.yml --output kotlin/gen/offer/user-offer-information && \
	../../emi.exe kotlin --path definitions/offer/offer-translations.emi.yml --output kotlin/gen/offer/offer-translations && \
	../../emi.exe kotlin --path definitions/offer/offer-management.emi.yml --output kotlin/gen/offer/offer-management

gen-golang:
	rm -rf golang/gen/ && \
	../../emi.exe go --path definitions/offer/user-offer-information.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/user-offer-information && \
	../../emi.exe go --path definitions/offer/offer-translations.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/offer-translations && \
	../../emi.exe go --path definitions/offer/offer-management.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/offer-management

gen-swift:
	rm -rf swift/gen/ && \
	../../emi.exe swift --path definitions/offer/user-offer-information.emi.yml --output swift/gen/offer/user-offer-information && \
	../../emi.exe swift --path definitions/offer/offer-translations.emi.yml --output swift/gen/offer/offer-translations && \
	../../emi.exe swift --path definitions/offer/offer-management.emi.yml --output swift/gen/offer/offer-management
