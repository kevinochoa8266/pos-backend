make test:
	go test ./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
	rm coverage.out

make build_container:
	docker build --tag sweetooth-backend .

make run_container: build_container
	docker run -d -p 8080:8080 --name pos-server sweetooth-backend

make delete_container:
	docker stop pos-server
	docker container rm pos-server
	docker rmi sweetooth-backend