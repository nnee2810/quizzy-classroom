.PHONY: migrate

deb:
	go get -u ./... && go mod tidy

run:
	go run main.go serve -e ./deploy/local.env

migrate:
	go run main.go migrate -e ./deploy/local.env
