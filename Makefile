
GOCMD=go
GOBINDATA=go-bindata
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

GOBUILD=$(GOCMD) build -v
GORUN=$(GOCMD) run -v

BUILD_DIR=$(PWD)/build

.PHONY: all
all: provider build

provider:
	$(GOBINDATA) -o configs/default_providers.go -pkg configs configs/default_providers.yaml

build:
	$(GOBUILD) -o $(BUILD_DIR)/terragraph .
	chmod +x $(BUILD_DIR)/terragraph

run:
	$(GORUN) .