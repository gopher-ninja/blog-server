IMAGE_TAG := latest
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BASE_DIR := $(shell pwd)
VERSION ?= $(shell git describe --exact-match 2> /dev/null || \
             git describe --match=$(git rev-parse --short=8 HEAD) \
             --always --dirty --abbrev=8)
BUILD_TGT := blog-server-$(VERSION)

.PHONY: all build blog-server docker clean 

all: build

build: blog-server

blog-server:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s -extldflags "-static"' -o $(BASE_DIR)/blog-server main.go


docker: build

	docker build . --build-arg=COMMIT=$(COMMIT) --build-arg=BRANCH=$(BRANCH) -t gopher-ninja/blog-server:$(IMAGE_TAG)
clean:
	rm -rf blog-server 

version:
	@echo ${VERSION}
