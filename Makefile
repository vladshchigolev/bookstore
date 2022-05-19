.PHONY: build
up_db:
	docker build --tag docker-mysql .
	docker run --name dockmysql -p 3306:3306 docker-mysql
run_server:
	go mod download
	go build -v ./main.go
	./main
test:
	go test bookstore/server
.DEFAULT_GOAL := build