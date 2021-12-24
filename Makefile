check-fmt:
	go fmt -n ./...

build-linux-amd64: check-fmt
	GOOS=linux GOARCH=amd64 go build -o ./out github.com/ymylei/gin-testing

test: check-fmt
	go test -v ./...
