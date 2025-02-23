DEFAULT: dev
.PHONY: help, up, down, dev, tidy

help:
	@echo "Makefile Help"
	@echo ""
	@echo "Available targets:"
	@echo "  make up      - Stops running containers (if any), builds the Docker images, and starts containers in detached mode."
	@echo "  make down    - Stops and removes the running Docker containers."
	@echo "  make dev     - Runs the application in development mode with debugging enabled."
	@echo "  make tidy    - Removes unused dependencies from the Go module."
	@echo "  make help    - Displays this help message."
	@echo ""
	@echo "Required Environment Variables:"
	@echo "  CONFIG_FILE         - Path to the configuration file (default: $(CONFIG_FILE))"
	@echo "  POSTGRES_USERNAME   - PostgreSQL username (default: $(POSTGRES_USERNAME))"
	@echo "  POSTGRES_PASSWORD   - PostgreSQL password (default: $(POSTGRES_PASSWORD))"


up:
	$(MAKE) down

	cd docker &&\
	docker compose build &&\
	docker compose up -d


down:
	cd docker &&\
	docker compose down

dev:
	DEBUG=true \
	CONFIG_FILE=/home/eurico-martins/Repos/Cerberus/example.env \
	POSTGRES_USERNAME=admin POSTGRES_PASSWORD=password \
	go run cmd/app/main.go


tidy:
	go mod tidy


