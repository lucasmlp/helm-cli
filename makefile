include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

add-repo-bitnami:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli repo-add bitnami https://charts.bitnami.com/bitnami

add-repo-stable:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli repo-add stable https://charts.helm.sh/stable

add-repo-local:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli repo-add localCharts '/Users/$(USER)/development/helm-charts'

add-k8s-api:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli add k8s-api

add-mysql:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli add mysql

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

install-mysql:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli install chart mysql mysql-dev

install-k8s-api:
	@ rm -rf helm-cli
	@ go build -o helm-cli ./cmd/main.go
	@ ./helm-cli install chart k8s-api k8s-api-dev