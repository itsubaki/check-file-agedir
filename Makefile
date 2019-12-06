.PHONY: build
build:
	rm -rf build
	mkdir build

	GOOS=linux GOARCH=amd64 go build
	zip build/check-file-agedir_linux_amd64.zip check-file-agedir

	GOOS=linux GOARCH=386 go build
	zip build/check-file-agedir_linux_386.zip check-file-agedir

	GOOS=darwin GOARCH=amd64 go build
	zip build/check-file-agedir_darwin_amd64.zip check-file-agedir
