.PHONY: migrate-fresh-supabase
migrate-fresh-supabase:
	migrate -path db/migrations -database "postgresql://postgres.settumozapjmoshlvqgf:9dGn99bPyoTVRBP5@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=disable" down
	migrate -path db/migrations -database "postgresql://postgres.settumozapjmoshlvqgf:9dGn99bPyoTVRBP5@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=disable" up

.PHONY: migrate-fresh
migrate-fresh:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/db_depublic?sslmode=disable" down
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/db_depublic?sslmode=disable" up

.PHONY: migrate
migrate:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/db_depublic?sslmode=disable" up 1

.PHONY: run-redis-wsl
run-redis-wsl:
	wsl --exec sudo service redis-server start

.PHONY: stop-redis-wsl
stop-redis-wsl:
	wsl --exec sudo service redis-server stop

.PHONY: jwt-key-generate
jwt-key-generate:
	go run .\build\generator\jwt\jwt_secret_key.go