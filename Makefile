run:
	go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

test:
	@go test ./... -short -cover
