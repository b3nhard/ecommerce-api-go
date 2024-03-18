build:
	go build  -o cmd/api cmd/api.go
run:build
	./cmd/api
