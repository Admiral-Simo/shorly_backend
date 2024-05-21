build:
	@CGO_ENABLED=1 GOOS="darwin" GOARCH="amd64" go build -o ./bin/api main.go
run: build
	@./bin/api
