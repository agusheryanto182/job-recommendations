# Development commands
dev:
	sudo docker-compose -f docker-compose.dev.yaml up --build

dev-down:
	sudo docker-compose -f docker-compose.dev.yaml down

# Production commands
prod:
	sudo docker-compose up --build

prod-down:
	sudo docker-compose down

prod-restart:
	sudo docker-compose down
	sudo docker-compose up --build

# Additional useful commands
logs:
	sudo docker-compose -f docker-compose.dev.yaml logs -f

ps:
	sudo docker-compose -f docker-compose.dev.yaml ps

clean:
	sudo docker-compose -f docker-compose.yml down -v --remove-orphans

.PHONY: dev dev-down prod prod-down logs ps clean