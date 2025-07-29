DB_URL=
migrate-up:
	migrate -path src/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path src/migrations -database "$(DB_URL)" down

create:
	migrate create -ext sql -dir src/migrations -seq $(name)
