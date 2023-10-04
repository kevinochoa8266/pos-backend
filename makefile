make test:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
	rm coverage.out

make build_container:
	docker build --tag sweetooth-backend .

make run_container:
	docker run -d -p 8080:8080 --name pos-server sweetooth-backend