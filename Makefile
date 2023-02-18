run:
	go run main.go

docker:
	docker-compose up -d

down:
	docker-compose down