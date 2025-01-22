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
	go run cmd/main.go