PKGS := $(shell go list ./... | grep -v vendor)

build: test
	go build

.PHONY: run
run: build
	./metapods

.PHONY: test
test:
	go test $(PKGS)

.PHONY: create-migration
migration:
	migrate create -dir db/migrations -ext sql $$NAME