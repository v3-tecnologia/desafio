.PHONY: build clean deploy lint mock


clean:
	rm -rf ./pkg/*

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o pkg/gps ./functions/gps/route.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o pkg/gyroscope ./functions/gyroscope/route.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o pkg/llm ./functions/llm/route.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o pkg/photo ./functions/photo/route.go

lint:
	go vet ./...

mock:
	 docker run -v "$PWD":/src -w /src vektra/mockery --all --keeptree --output /internal/test/mock

test:
	go test -count=1 ./internal/usecase/...