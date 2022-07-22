.PHONY: build-cli
build-cli:
	go build -v -o ./bin/oras-tools ./cmd/cli
	chmod +x ./bin/oras-tools
