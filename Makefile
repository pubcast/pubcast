PKGS := $(shell go list ./... | grep -v vendor)

build: test
	go build

.PHONY: run
run: build
	./metapods

.PHONY: test
test:
	go test $(PKGS)

.PHONY: migration
migration:
	migrate create -dir data/migrations -ext sql $$NAME

.PHONY: database
database:
	./scripts/create_db.sh

.PHONY: migrate-up
migrate-up: database
	# We use ?sslmode=disable to accommodate for crappy brew installs
	migrate -source file://data/migrations -database postgres://localhost:5432/metapods?sslmode=disable up
	@echo "✨ Finished."

.PHONY: drop-database
drop-database:
	psql -U postgres -c "drop database metapods"
	psql -U postgres -c "drop database metapods_test"