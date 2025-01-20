.PHONY: up-local

# Container Commands

up-local:
	$(MAKE) down-local

	cd docker &&\
	docker compose build &&\
	docker compose up -d

down-local:
	cd docker &&\
	docker compose down