.PHONY: all metadata movie rating consul

services: metadata movie rating

consul:
	docker run -d --rm --name=dev-consul -p 8500:8500 -p 8600:8600/udp hashicorp/consul:1.20 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

metadata:
	go run ./metadata/cmd/main.go

movie:
	go run ./movie/cmd/main.go

rating:
	go run ./rating/cmd/main.go

