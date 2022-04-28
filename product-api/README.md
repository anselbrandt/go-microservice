# Generate `swagger.yaml` and Client

Ensure `swagger.yaml` is up to date by running `make swagger` in server folder.

Then run `make generate_client`, which will output:

```
For this generation to compile you need to have some packages in your GOPATH:

	* github.com/go-openapi/errors
	* github.com/go-openapi/runtime
	* github.com/go-openapi/runtime/client
	* github.com/go-openapi/strfmt

You can get these now with: go get -u -f ./...
```
