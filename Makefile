generate:
	protoc --go_out=. --go-grpc_out=. api/user.proto

clean:
	rm -rf pkg

run:
	cd cmd; go run main.go

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
