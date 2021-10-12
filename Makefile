generate:
	cd api; buf generate

clean:
	rm -rf pkg

run:
	cd cmd; go run main.go
