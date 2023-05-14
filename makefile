include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

add-local:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli add k8s-api

add-ark:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli add ark

index:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli index

images:
	@ rm -rf helm-cli
	@ rm -rf index.yaml
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli images

build:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go

install-ark:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli install ark

install-k8s-api:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli install k8s-api