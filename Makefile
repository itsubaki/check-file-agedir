.PHONY: build
build:
	rm -rf build
	mkdir build

	GOOS=linux GOARGH=amd64 go build
	zip build/check-fileagedir_linux_amd64.zip check-fileagedir

	GOOS=linux GOARGH=386 go build
	zip build/check-fileagedir_linux_386.zip check-fileagedir

	GOOS=darwin GOARGH=amd64 go build
	zip build/check-fileagedir_darwin_amd64.zip check-fileagedir
