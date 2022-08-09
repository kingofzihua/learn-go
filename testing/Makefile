
.PHONY: test

test:
	 go test -v ./... -coverprofile coverage.out
	 go tool cover -func coverage.out
	 go tool cover -html=coverage.out -o index.html

.DEFAULT_GOAL := test
