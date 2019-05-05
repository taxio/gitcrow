# gitcrow

## Set up DB
Dependencies
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- PostgreSQL

Commands
- `make start-db`: create and start
- `make stop-db`: stop db container
- `make rm-db`: remove db container and volume

### Migration
TODO: migrate by script

At the first of migration, you must create database and schema on your db.
1. login db using `psql`
1. `create database gitcrow;`
1. `\c gitcrow`
1. `create schema gitcrow;`

migration using [golang-migrate/migrate](https://github.com/golang-migrate/migrate).
```bash
> migrate -database 'postgresql://{{Domain or IP}}:{{PORT}}/gitcrow?user={{USERNAME}}&password={{PASSWORD}}&sslmode=disable' -path ./migrations/ up
```
