run:
	go run main.go

up:
	sudo docker-compose up -d

start:
	sudo docker-compose up

lint:
	goliangci-lint run .

restart:
	sudo docker-compose restart

stop:
	sudo docker-compose stop

exec:
	sudo docker-compose exec nlp

ps:
	sudo docker-compose ps
