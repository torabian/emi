gen:
	make gen-kotlin && make gen-js && make gen-ts && make gen-golang

gen-js:
	rm -rf javascript/gen/ && \
	emi js --path definitions/offer/user-offer-information.emi.yml --output javascript/gen/offer/user-offer-information && \
	emi js --path definitions/offer/offer-translations.emi.yml --output javascript/gen/offer/offer-translations && \
	emi js --path definitions/offer/offer-management.emi.yml --output javascript/gen/offer/offer-management

gen-ts:
	rm -rf typescript/gen/ && \
	emi js --path definitions/offer/user-offer-information.emi.yml --output typescript/gen/offer/user-offer-information --tags typescript && \
	emi js --path definitions/offer/offer-translations.emi.yml --output typescript/gen/offer/offer-translations --tags typescript && \
	emi js --path definitions/offer/offer-management.emi.yml --output typescript/gen/offer/offer-management --tags typescript

gen-kotlin:
	rm -rf kotlin/gen/ && \
	emi kotlin --path definitions/offer/user-offer-information.emi.yml --output kotlin/gen/offer/user-offer-information && \
	emi kotlin --path definitions/offer/offer-translations.emi.yml --output kotlin/gen/offer/offer-translations && \
	emi kotlin --path definitions/offer/offer-management.emi.yml --output kotlin/gen/offer/offer-management

gen-golang:
	rm -rf golang/gen/ && \
	emi go --path definitions/offer/user-offer-information.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/user-offer-information && \
	emi go --path definitions/offer/offer-translations.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/offer-translations && \
	emi go --path definitions/offer/offer-management.emi.yml --emigo github.com/torabian/emi/public/allegro-sdk/golang/emigo --output golang/gen/offer/offer-management

gen-swift:
	rm -rf swift/gen/ && \
	emi swift --path definitions/offer/user-offer-information.emi.yml --output swift/gen/offer/user-offer-information && \
	emi swift --path definitions/offer/offer-translations.emi.yml --output swift/gen/offer/offer-translations && \
	emi swift --path definitions/offer/offer-management.emi.yml --output swift/gen/offer/offer-management
