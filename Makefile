run:
	go run main.go
docker:
	docker-compose up -d
lint:
	goliangci-lint run .


