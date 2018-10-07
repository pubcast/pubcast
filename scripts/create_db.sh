#!/bin/bash

# Regular DB
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'pubcast'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE pubcast"

# Testing DB
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'pubcast_test'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE pubcast_test"
