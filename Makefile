deb:
	go get -u ./... && go mod tidy

run-local:
	go run main.go serve -e ./deploy/local.env

migration-local:
	go run main.go migration -e ./deploy/local.env
