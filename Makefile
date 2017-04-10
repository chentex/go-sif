PKGS = $$(go list ./... | grep -v /vendor/)


all: get-deps test

get-deps: install-glide dependencies

install-glide:
	mkdir -p $(GOPATH)/bin
	curl https://glide.sh/get | sh

dependencies:
	glide install

test:
	go test $(PKGS)

.PHONY: get-deps install-glide dependencies test