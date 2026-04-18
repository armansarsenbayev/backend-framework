DB_URL=postgresql://postgres:1234@localhost:5432/food_delivery_v2?sslmode=disable
MIGRATE=C:\Users\sabal\go\bin\migrate

migrateup:
	"$(MIGRATE)" -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	"$(MIGRATE)" -path db/migration -database "$(DB_URL)" -verbose down

migratecreate:
	"$(MIGRATE)" create -ext sql -dir db/migration -seq $(name)

.PHONY: migrateup migratedown migratecreate
