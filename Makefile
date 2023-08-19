run:
	go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

mocks:
	@mockery --all --keeptree --dir=modules --output=utilities/mocks --case underscore

test:
	@go test ./... -short -cover

test-cover:
	@echo "\x1b[32;1m>>> running unit test and calculate coverage\x1b[0m"
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./... | grep -v mocks | tr '\n' ',')
	@go tool cover -func=coverage.txt