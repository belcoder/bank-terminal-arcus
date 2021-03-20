.SILENT:

build:
	env GOOS=linux GOARCH=amd64 GOARM=7 go build -o .bin/bank-terminal-arcus cmd/app/main.go
build-emulation:
	env GOOS=linux GOARCH=amd64 GOARM=7 go build -o .data/arcus/terminal-emulation cmd/emulation/main.go
