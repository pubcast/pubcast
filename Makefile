PKGS := $(shell go list ./... | grep -v vendor)

# Builds a new pubcast binary
build: test
	go build

# Runs a local pubcast instance
.PHONY: run
run: build
	./pubcast

# Runs the tests
.PHONY: test
test:
	go test $(PKGS) -cover

# Installs go dep
$(GOPATH)/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

# ----- Database ----- # 

# Installs our db migration tool
$(GOPATH)/bin/migrate: $(GOPATH)/bin/dep
	@go get -u github.com/lib/pq
	@go get -u github.com/golang-migrate/migrate/cli
	@cd $(GOPATH)/src/github.com/golang-migrate/migrate/cli
	@dep ensure
	@go build -tags 'postgres' -o $(GOPATH)/bin/migrate github.com/golang-migrate/migrate/cli

# Creates a new sql migration file in ./data/migrations
.PHONY: migration
migration:
	migrate create -dir data/migrations -ext sql $$NAME

# Creates a test and local database 
.PHONY: database
database:
	./scripts/create_db.sh

# Migrates the database up to the newest version
.PHONY: migrate-up
migrate-up: database $(GOPATH)/bin/migrate
	# We use ?sslmode=disable to accommodate for crappy brew installs
	migrate -source file://data/migrations -database postgres://localhost:5432/pubcast?sslmode=disable up
	migrate -source file://data/migrations -database postgres://localhost:5432/pubcast_test?sslmode=disable up
	@echo "✨ Finished."

# Removes the current connected databases
.PHONY: drop-database
drop-database:
	psql -U postgres -c "drop database pubcast"
	psql -U postgres -c "drop database pubcast_test"

# ----- Docker ----- #

NAME := pubcast/pubcast
TAG := $(shell git log -1 --pretty=%h)
IMG := ${NAME}:${TAG}
LATEST := ${NAME}:latest

# Makes a new docker image
image: 
	docker build -t ${IMG} .
	docker tag ${IMG} ${LATEST}

# Pushes an existing image to dockerhub
push-image:
	docker push ${NAME}

# Logs us in to dockerhub
# Requires the env variables DOCKER_USER and DOCKER_PASS
docker-login:
	docker log -u ${DOCKER_USER} -p ${DOCKER_PASS}