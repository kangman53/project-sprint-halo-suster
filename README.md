# Project Sprint Halo Suster App

HaloSus is a backend that nurses can use to record patient medical records.
# How to run local (dev purposes)

- Create file `halo-suster.env`

```go
DB_NAME=
DB_HOST=
DB_USERNAME=
DB_PASSWORD=
DB_PORT=5432

BCRYPT_SALT=8
JWT_SECRET=
```

- run `make build-dev`
- run `make run-dev`

if you're running this for the first time, do:
- run `make migrate-db`


# Stacks
- Golang >1.21.0
- Go Fiber
- Postgres
- Docker
