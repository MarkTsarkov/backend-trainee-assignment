.PHONY: dc run test lint

dc:
	docker-compose up --remove-orphans --build

run:
	go build -o app cmd/avito/main.go && HTTP_ADDR=:8080 ./app

