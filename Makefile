init: docker-down-clear \
	docker-pull docker-build docker-up \
	app-init

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull
docker-build:
	docker-compose build --pull
docker-up:
	docker-compose up -d

app-init: run migrate-up migrate-up-ch

generate:
	protoc --go_out=. --go-grpc_out=. api/user.proto

clean:
	rm -rf pkg

run:
	cd cmd; go run main.go

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate-up-ch:
	migrate -path ./schema-ch -database 'clickhouse://localhost:9000?username=default&database=default' up

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-create-ch:
	migrate create -ext sql -dir ./schema-ch -seq init

