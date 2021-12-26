check:
	goimports -l -d .
	staticcheck ./...

build-default: check
	go build -o ./out/gin-testing github.com/ymylei/gin-testing

build-linux-arm64: check
	GOOS=linux GOARCH=arm64 go build -o ./out/gin-testing-linux-aarch64 github.com/ymylei/gin-testing

build-linux-amd64: check
	GOOS=linux GOARCH=amd64 go build -o ./out/gin-testing-linux-amd64 github.com/ymylei/gin-testing

build-docker: build-linux-arm64
	docker build . -t gin-testing:latest

test: check
	go test -v ./...

format:
	goimports -l -w .
