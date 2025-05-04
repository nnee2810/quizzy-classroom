.PHONY: http grpc migrate

deb:
	go get -u ./... && go mod tidy

http:
	go run main.go http -e ./deploy/local.env

grpc:
	go run main.go grpc -e ./deploy/local.env

migrate:
	go run main.go migrate -e ./deploy/local.env
