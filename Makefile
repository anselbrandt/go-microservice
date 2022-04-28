check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	cd product-api && swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	cd product-api && mkdir sdk && cd sdk && swagger generate client -f ../swagger.yaml -A product-api && go get -u -f ./...