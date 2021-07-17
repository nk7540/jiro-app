#!/bin/sh

sql-migrate down
sql-migrate up
sqlboiler mysql

go run src/cmd/seed/seed.go
