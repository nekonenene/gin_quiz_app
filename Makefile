.PHONY: init
init:
	dep ensure
	cp default.env .env

.PHONY: db.migrate
db.migrate:
	docker-compose exec api /bin/bash -c 'mysqldef -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: db.migrate.dry_run
db.migrate.dry_run:
	docker-compose exec api /bin/bash -c 'mysqldef --dry-run -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db --file=schema.sql $$MYSQL_DATABASE'

.PHONY: db.export
db.export:
	docker-compose exec api /bin/bash -c 'mysqldef -u $$MYSQL_USER -p $$MYSQL_PASSWORD -h db $$MYSQL_DATABASE --export'

.PHONY: mysql
mysql:
	docker-compose exec db /bin/bash -c 'mysql -u $$MYSQL_USER --password=$$MYSQL_PASSWORD $$MYSQL_DATABASE'
