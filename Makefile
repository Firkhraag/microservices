swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

run:
	go run cmd/microservices/main.go

.DEFAULT_GOAL := run