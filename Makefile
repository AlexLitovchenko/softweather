.SILENT:

build:
	go build -o myapp ./cmd

compose-up: build
	docker-compose up --build 

compose-stop:
	docker-compose stop

run-cli: 
	cd cli && go run main.go ${STRING} http://localhost:8080/api/substring
	