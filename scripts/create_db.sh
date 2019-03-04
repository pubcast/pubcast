#!/bin/bash

# Fetch Postgres image from docker
if [ -z $(docker images -q postgres:11) ]; then
    docker pull postgres:11
fi

# Run Postgres image
docker run --rm --name pubcast-db -e POSTGRES_USER=postgres -d -p 5432:5432 postgres

# Regular DB
psql --host=localhost -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'pubcast'" | grep -q 1 || psql --host=localhost -U postgres -c "CREATE DATABASE pubcast"

# Testing DB
psql --host=localhost -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'pubcast_test'" | grep -q 1 || psql --host=localhost -U postgres -c "CREATE DATABASE pubcast_test"
