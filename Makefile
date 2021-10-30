generate:
	protoc --go_out=. --go-grpc_out=. api/user.proto

clean:
	rm -rf pkg

run:
	cd cmd; go run main.go
