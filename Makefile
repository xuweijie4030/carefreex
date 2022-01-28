# const
GO_CMD=go
GO_BUILD=$(GO_CMD) build
SERVICE_NAME=carefreex
REGISTRY=etcdV3
OUTPUT_DIR=./bin/$(SERVICE_NAME)

default: build

build:
	$(GO_BUILD) -tags "$(REGISTRY)" -o $(OUTPUT_DIR)

build-win:
	$(GO_BUILD) -tags "$(REGISTRY)" -o $(OUTPUT_DIR).exe

clean:
	rm -rf $(OUTPUT_DIR)
