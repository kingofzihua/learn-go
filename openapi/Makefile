include .env

.PHONY: openapi_doc
openapi_doc:
	protoc  --proto_path=./proto \
		 	--proto_path=./third_party \
			--openapi_out=title=${PROJECT_NAME},version=${PROJECT_VERSION},description=${PROJECT_DESCRIPTION}:./openapi \
			proto/*.proto

.PHONY: openapi_doc_clean
openapi_doc_clean:
	rm -r ./openapi/*.yaml
	
help:
	@echo 'help'