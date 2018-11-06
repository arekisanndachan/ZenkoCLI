GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BIN_OUTPUT_DIR=$(CURDIR)/$(bin)

VENDOR_AWS_DIR=$(CURDIR)/vendor/aws/aws-sdk-go
GEN_API_DIR=$(VENDOR_AWS_DIR)/private/model/cli/gen-api

install-deps:
	@echo Installing dep
	$(GOGET) -v github.com/golang/dep/cmd/dep

build-gen:
	$(GOBUILD) -o $(BIN_OUTPUT_DIR)/gen-api $(GEN_API_DIR)
