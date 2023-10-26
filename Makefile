.PHONY: all
all: setup build

.PHONY: setup
setup:
	go mod download
	go mod tidy

.PHONY: build
build: dist clean
	go build -o ./dist/toxiproxy-server ./cmd/server.go
	go build -o ./dist/toxiproxy-cli github.com/Shopify/toxiproxy/v2/cmd/cli

dist:
	mkdir -p dist

.PHONY: clean
clean:
	rm -fr dist/*