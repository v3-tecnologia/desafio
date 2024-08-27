build-mocks:
	@go get github.com/vektra/mockery/v2@latest 
	@go install github.com/vektra/mockery/v2@latest 
	@~/go/bin/mockery --dir ./service --output ./service/mocks --all

tests: 
	go test -v -cover ./...
