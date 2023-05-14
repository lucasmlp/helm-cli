include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

add-local:
	@ rm -rf main
	@ go build -o helm-cli ./cmd/main.go
	@ ./main add k8s-api

add-web:
	@ rm -rf main
	@ go build -o helm-cli ./cmd/main.go
	@ ./main add ark

index:
	@ rm -rf main
	@ go build -o helm-cli ./cmd/main.go
	@ ./main index

images:
	@ rm -rf main
	@ rm -rf index.yaml
	@ go build -o helm-cli ./cmd/main.go
	@ ./main images

build:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go