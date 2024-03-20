include .env
DSN = mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):3306)/$(DB_NAME)?parseTime=true

dev/up:
	docker-compose up -d api database
dev/down:
	docker-compose down
dev/build:
	docker-compose build

dev/migrate/up:
	docker-compose run --rm migrate migrate -path ./migrations -database "$(DSN)" up
dev/migrate/create:
	docker-compose run --rm migrate migrate create -ext sql -dir /app/migrations $(name)