DOCKER_IMAGE ?= $(shell basename `git rev-parse --show-toplevel`)

.PHONY: all
all: build

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Build

.PHONY: build
build: fmt vet ## Run go build against code.
	go build ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test $$(go list ./... | grep -v /e2e) -coverprofile cover.out


.PHONY: build-docker
build-docker: ## Build the docker image.
	docker build -t $(DOCKER_IMAGE):latest .