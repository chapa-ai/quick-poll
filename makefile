run:
	docker-compose build && docker-compose up

tidy:
	go mod tidy