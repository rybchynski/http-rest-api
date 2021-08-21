.PHONY: build
build: 
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run: db
	go run ./cmd/apiserver

.PHONY: stop
stop:
	docker-compose stop

.PHONY: db
db:
	docker-compose up -d

.DEFAULT_GOAL := build
