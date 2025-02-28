.PHONY: build
build:
	docker-compose build $(OPTIONS)

.PHONY: up
up: 
	$(MAKE) build
	docker-compose up -d

.PHONY: down
down:
	docker-compose down $(OPTIONS)

.PHONY: sql
sql:
	docker-compose exec db bash -c 'mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'

.PHONY: migrate
migrate:
	docker-compose exec backend sh -c 'go run migrate/migrate.go'

.PHONY: up-all
up-all:
	$(MAKE) up
	$(MAKE) migrate