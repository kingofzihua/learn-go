include .env

.PHONY: openapi_doc_clean
# clean openapi doc
openapi_doc_clean:
	rm -r ./openapi/*.yaml

.PHONY: openapi_doc
# proto to openapi doc
openapi_doc:
	protoc  --proto_path=./proto \
		 	--proto_path=./third_party \
			--openapi_out=title=${PROJECT_NAME},version=${PROJECT_VERSION},description=${PROJECT_DESCRIPTION}:./openapi \
			proto/category.proto

.PHONY: openapi_js
# openapi/openapi.yaml to javascript code
openapi_js:
	docker run --rm --env "JAVA_OPTS=-Dlog.level=error" -v "${PWD}:/local" --platform linux/amd64 \
		"openapitools/openapi-generator-cli:v5.4.0" generate \
		-i "/local/openapi/openapi.yaml" \
		-g javascript \
		-o "/local/web/src/openapi"
	
.PHONY: oapi_codegen
# show oapi-codegen
oapi_codegen:
	oapi-codegen openapi/openapi.yaml > oapi-codegen/openapi.gen.go

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

