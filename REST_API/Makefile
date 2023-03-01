
//Делаем для того, чтобы записывать код. 

.PHONY: build 
build:
go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ... 


//Условия по умолчанию 

.DEFAULT_GOAL := build