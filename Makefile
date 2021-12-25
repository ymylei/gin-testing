check-fmt:
	go fmt -n ./...

build-default: check-fmt
	go build -o ./out/gin-testing github.com/ymylei/gin-testing

build-linux-arm64: check-fmt
	GOOS=linux GOARCH=arm64 go build -o ./out/gin-testing-linux-arm64 github.com/ymylei/gin-testing

build-linux-amd64: check-fmt
	GOOS=linux GOARCH=amd64 go build -o ./out/gin-testing-linux-amd64 github.com/ymylei/gin-testing

build-docker: build-linux-arm64
	docker build . -t gin-testing:latest

test: check-fmt
	go test -v ./...
