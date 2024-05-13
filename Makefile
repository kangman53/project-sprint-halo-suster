build-dev:
	docker compose build

restart-dev:
	docker restart halosus-web

run-dev:
	docker compose up -d

logs-web:
	docker logs -f --tail 100 halosus-web

logs-db:
	docker logs -f halosus-db

check-db:
	docker exec -it halosus-db psql -U halosus -d halosus-db

clear-db:
	docker rm -f -v halosus-db

migrate-db:
	migrate -database "postgres://halosus:password@localhost:5432/halosus-db?sslmode=disable" -path database/migrations up
	
migrate-db-down:
	migrate -database "postgres://halosus:password@localhost:5432/halosus-db?sslmode=disable" -path database/migrations down -all
	
build-prod-linux:
	GOOS=linux GOARCH=amd64 go build -o build/halo-suster

build-prod-win:
	GOOS=windows GOARCH=amd64 go build -o build/halo-suster.exe

build-prod-mac:
	GOOS=darwin GOARCH=amd64 go build -o build/halo-suster

build-prod-docker:
	docker build . -t halo-suster
	docker tag halo-suster:latest rereasdev/halo-suster:latest

docker-push:
	docker push rereasdev/halo-suster:latest
