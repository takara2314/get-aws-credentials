.PHONY: build
build:
	go build -o get-aws-credentials -ldflags="-s -w" -trimpath main.go
