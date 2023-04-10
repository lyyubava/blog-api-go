#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
	CREATE USER apiuser with password 'example';
	CREATE DATABASE blog;
	GRANT ALL PRIVILEGES ON DATABASE blog TO apiuser;
EOSQL