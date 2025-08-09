build:
	@go build -o bin/go-http-template cmd/main.go

run: build
	@./bin/go-http-template

migrate-create:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/script/main.go up

migrate-down:
	@go run cmd/script/main.go down