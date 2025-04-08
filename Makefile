deb:
	go get -u ./... && go mod tidy

run:
	go run main.go serve -e ./deploy/local.env

migration:
	go run main.go migration -e ./deploy/local.env
