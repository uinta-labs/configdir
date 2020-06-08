
example_main = $(PWD)/examples/print_dirs.go

build-example-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/example-darwin-amd64 $(example_main)

build-example-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/example-linux-amd64 $(example_main)

build-example-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/example-windows-amd64 $(example_main)

all: build-example-darwin build-example-linux build-example-windows

clean:
	rm -rf example/

.PHONY: build-example-darwin build-example-linux build-example-windows
