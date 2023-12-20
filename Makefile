GOPATH:=$(shell go env GOPATH)

build_linux: 
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags="linux" -o test-amd64 test.go

build_mac: 
	go build -tags="mac" -o test-mac test.go