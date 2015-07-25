default: build

build:
	`<.env xargs -I % echo "export %"`; \
	go build -o hoard-server *.go

run:
	`<.env xargs -I % echo "export %"`; \
	go run *.go
