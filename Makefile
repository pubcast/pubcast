PKGS := $(shell go list ./... | grep -v vendor)

build: test
	go build

.PHONY: run
run: build
	./metapods

.PHONY: test
test:
	go test $(PKGS)

$(GOPATH)/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

$(GOPATH)/bin/migrate: $(GOPATH)/bin/dep
	@go get -u github.com/golang-migrate/migrate/cli
	@cd $(GOPATH)/src/github.com/golang-migrate/migrate/cli
	@dep ensure
	@go build -tags 'postgres' -o $(GOPATH)/bin/migrate github.com/golang-migrate/migrate/cli

.PHONY: migration
migration:
	migrate create -dir data/migrations -ext sql $$NAME

.PHONY: database
database:
	./scripts/create_db.sh

.PHONY: migrate-up
migrate-up: database $(GOPATH)/bin/migrate
	# We use ?sslmode=disable to accommodate for crappy brew installs
	migrate -source file://data/migrations -database postgres://localhost:5432/metapods?sslmode=disable up
	migrate -source file://data/migrations -database postgres://localhost:5432/metapods_test?sslmode=disable up
	@echo "✨ Finished."

.PHONY: drop-database
drop-database:
	psql -U postgres -c "drop database metapods"
	psql -U postgres -c "drop database metapods_test"