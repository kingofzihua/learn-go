include .env

.PHONY: openapi_doc_clean
openapi_doc_clean:
	rm -r ./openapi/*.yaml

.PHONY: openapi_doc
openapi_doc:
	protoc  --proto_path=./proto \
		 	--proto_path=./third_party \
			--openapi_out=title=${PROJECT_NAME},version=${PROJECT_VERSION},description=${PROJECT_DESCRIPTION}:./openapi \
			proto/*.proto

.PHONY: openapi_js
openapi_js:
	docker run --rm --env "JAVA_OPTS=-Dlog.level=error" -v "${PWD}:/local" --platform linux/amd64 \
		"openapitools/openapi-generator-cli:v5.4.0" generate \
		-i "/local/openapi/openapi.yaml" \
		-g typescript-axios \
		-o "/local/web/src/openapi"
	
help:
	@echo 'help'