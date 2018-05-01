.PHONY: build
build:
	rm -rf build
	mkdir build

	GOOS=linux GOARGH=amd64 go build
	zip build/check-file-agedir_linux_amd64.zip check-file-agedir

	GOOS=linux GOARGH=386 go build
	zip build/check-file-agedir_linux_386.zip check-file-agedir

	GOOS=darwin GOARGH=amd64 go build
	zip build/check-file-agedir_darwin_amd64.zip check-file-agedir
