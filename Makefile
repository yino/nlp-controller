run:
	go run main.go

docker:
	sudo docker-compose up -d

docker-test:
	sudo docker-compose up

lint:
	goliangci-lint run .

docker-restart:
	sudo docker-compose restart

