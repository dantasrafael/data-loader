# Makefile
run:
	go run main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s" -o data-loader

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s" -o data-loader

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s" -o data-loader

update-module:
	go mod tidy
