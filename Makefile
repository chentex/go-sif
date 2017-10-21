# Meta info
NAME := gosif
VERSION := v0.2.1
MAINTAINER := "Vicente Zepeda <chente.z.m@gmail.com"
# LD Flags
DATE := $(shell date -u +%Y%m%d.%H%M%S)
COMMIT_ID := $(shell git rev-parse --short HEAD)
GIT_REPO := $(shell git config --get remote.origin.url)
# Go tools flags
LD_FLAGS := -X github.com/chentex/go-sif/cmd.buildVersion=$(VERSION)
LD_FLAGS += -X github.com/chentex/go-sif/cmd.buildCommit=$(COMMIT_ID)
LD_FLAGS += -X github.com/chentex/go-sif/cmd.buildDate=$(DATE)
EXTRA_BUILD_VARS := CGO_ENABLED=0 GOARCH=amd64
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)


all: install_dependencies test package-linux package-darwin

lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)

test: lint
	 @go test -v $(SOURCE_DIRS) -cover -bench . -race

install_dependencies:
	glide install

cover:
	@bash cover.sh

binaries: binary-darwin binary-linux

binary-darwin:
	@-rm -rf build/dist/darwin
	@-mkdir -p build/dist/darwin
	GOOS=darwin $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/darwin/$(NAME)

binary-linux:
	@-rm -rf build/dist/linux
	@-mkdir -p build/dist/linux
	GOOS=linux $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/linux/$(NAME)


package-darwin: binary-darwin
	@tar -czf build/dist/$(NAME).darwin-amd64.tar.gz -C build/dist/darwin $(NAME)


package-linux: binary-linux
	@tar -czf build/dist/$(NAME).linux-amd64.tar.gz -C build/dist/linux $(NAME)

.PHONY: lint test install_dependencies cover binaries binary-darwin binary-linux package-darwin package-linux