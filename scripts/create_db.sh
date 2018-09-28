#!/bin/bash

# Regular DB
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'metapods'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE metapods"

# Testing DB
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'metapods_test'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE metapods_test"
