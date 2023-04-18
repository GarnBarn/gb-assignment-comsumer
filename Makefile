run-gb-assignment-create-consumer:
	ENV=local go run cmd/gb-assignment-create-consumer/main.go

run-compose:
	COMPOSE_PROJECT_NAME=garnbarn-assignment docker-compose up -d