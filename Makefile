.PHONY: init
init:
	go mod vendor
	cp default.env .env

.PHONY: db_migrate
db_migrate:
	docker-compose exec api /bin/bash -c 'mysqldef -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: db_migrate_dry
db_migrate_dry:
	docker-compose exec api /bin/bash -c 'mysqldef --dry-run -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: db_export
db_export:
	docker-compose exec api /bin/bash -c 'mysqldef -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db $$MYSQL_DATABASE --export'

.PHONY: mysql
mysql:
	docker-compose exec db /bin/bash -c 'mysql -u $$MYSQL_USER --password=$$MYSQL_PASSWORD $$MYSQL_DATABASE'
