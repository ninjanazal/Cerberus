DEFAULT: dev
.PHONY: up, down

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