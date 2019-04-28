.PHONY: init
init:
	dep ensure
	cp default.env .env

.PHONY: migrate
migrate:
	docker-compose exec api /bin/bash -c 'mysqldef -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: migrate_dry
migrate_dry:
	docker-compose exec api /bin/bash -c 'mysqldef --dry-run -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: mysql
mysql:
	docker-compose exec db /bin/bash -c 'mysql -u $$MYSQL_USER --password=$$MYSQL_PASSWORD $$MYSQL_DATABASE'
