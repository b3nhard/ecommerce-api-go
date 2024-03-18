build:
	go build -ldflags '-s -w' -o cmd/api cmd/api.go
run:build
	./cmd/api

test:
	go test
