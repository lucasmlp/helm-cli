GOPATH=$(shell go env GOPATH)

add-local:
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main add k8s-api

add-web:
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main add ark

index:
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main index

images:
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main images