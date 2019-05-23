PKGS = $(shell go list ./... | grep -v /test)

build-client:
	CGO_ENABLED=0 go build -o ./build/client ./cmd/client
.PHONY: build-client

build-server:
	CGO_ENABLED=0 go build -o ./build/server ./cmd/server
.PHONY: build-server

build: build-server build-client
.PHONY: build

lint:
	golint $(PKGS) 
.PHONY: lint

test-unit: # TODO: Please implement
	go test --race --cover -v $(PKGS)
.PHONY: test-unit

test-integration:
	go test --race -v test/integration_test.go
.PHONY: test-integration

test-benchmark:
	go test -v -bench=. test/benchmark_test.go
.PHONY: test-benchmark

test: lint test-unit test-integration
.PHONY: test
