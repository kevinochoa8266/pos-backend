make test:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
	rm coverage.out
