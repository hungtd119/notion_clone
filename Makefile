DB_URL=db

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

create:
	migrate create -ext sql -dir migrations -seq $(name)
