up_build: build
	sudo docker compose down
	sudo docker compose up --build

down:
	sudo docker compose down

build:
	env CGO_ENABLED=0 go build -o binaryApp ./cmd/api
