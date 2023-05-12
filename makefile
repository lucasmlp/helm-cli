GOPATH=$(shell go env GOPATH)

cli-local:
	@ echo
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main add k8s-api

cli-web:
	@ echo
	@ rm -rf main
	@ go build ./cmd/main.go
	@ ./main add ark