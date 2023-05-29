fmt:
	go fmt ./...

fmt-check:
	gofmt -l .
	@test -z "$(shell gofmt -l .)"

build:
	CGO_ENABLED=0 go build

build-all:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o dist/hello-linux-amd64
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o dist/hello-darwin-arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/hello-windows-amd64.exe

version-bump:
ifndef VERSION
	$(error VERSION is undefined)
endif
	slu go-code version-bump --version ${VERSION}
