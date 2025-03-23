help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  download - Download Go dependencies"
	@echo "  tidy - Tidy Go dependencies"
	@echo "  run - Run the application"
	@echo "  add_build_deps - Add build dependencies"
	@echo "  docker_up - Start the Docker containers"

download:
	go mod download

tidy:
	go mod tidy

run:
	source .env
	go build -o main .

add_build_deps:
	apk add --no-cache gcc musl-dev

docker_up:
	source .env
	docker compose up -d

docker_stop:
	docker stop backend-backend-1

docker_start:
	docker start backend-backend-1

