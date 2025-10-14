.dEFAULT_GOAL := build


.PHONY: fmt vet build


fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean

windows:
	GOOS=windows go build
