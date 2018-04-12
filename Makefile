SHELL := /bin/bash

mysqlenv:
	export DATABASE_URL="dbuser_f:dbpass_f@tcp(localhost:3306)/freelance?multiStatements=true"

pgsqlenv:
	export DATABASE_URL="postgres://dbuser_f:dbpass_f@localhost:5432/freelance?query"

