run:
	go run main.go
docker:
	docker-compose up -d
docker-test:
	docker-compose up
lint:
	goliangci-lint run .

docker-restart:
	docker-compose restart

