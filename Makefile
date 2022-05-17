USERNAME := alekseevdenis
APP_NAME := rdr
VERSION := latest
PROJECT := rdr
GIT_COMMIT := $(shell git rev-parse HEAD)

check:
	golangci-lint run -c golangci-lint.yaml

test:
	go test ./...

generate:
	go generate ./...

run:
	go install -ldflags="-X '$(PROJECT)/app/config.Version=$(VERSION)' \
	-X '$(PROJECT)/app/config.Commit=$(GIT_COMMIT)'" ./cmd/redirector && redirector

build_container:
	docker build --build-arg=GIT_COMMIT=$(GIT_COMMIT) --build-arg=VERSION=$(VERSION) --build-arg=PROJECT=$(PROJECT) \
	-t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) .
