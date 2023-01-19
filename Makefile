run:
	go run *.go

docker:
	docker-compose up -d

down:
	docker-compose down