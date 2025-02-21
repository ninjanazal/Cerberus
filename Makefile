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
	@echo "Notes:"
	@echo "- 'up' first stops any running containers by calling 'down' before rebuilding and starting new ones."
	@echo "- 'dev' expects a CONFIG_FILE environment variable to be set."

up:
	$(MAKE) down

	cd docker &&\
	docker compose build &&\
	docker compose up -d

down:
	cd docker &&\
	docker compose down

dev:
	DEBUG=true CONFIG_FILE=${CONFIG_FILE} go run cmd/app/main.go

# Remove all unused dependencies
tidy:
	go mod tidy


