# コンテナを起動
docker_up:
	docker compose -f docker-compose.local.yml up -d --build
	docker compose -f docker-compose.test.yml up -d --build

# wireを自動生成
docker_wire_gen:
	docker compose -f docker-compose.local.yml exec gen wire api/di/wire.go
	docker compose -f docker-compose.local.yml exec gen wire auth/di/wire.go
	docker compose -f docker-compose.local.yml exec gen wire batch/di/wire.go

# 全てのappを自動生成
docker_app_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/request/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./api/presentation/request...
	docker compose -f docker-compose.local.yml exec gen go generate ./script/response/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./api/presentation/response...

# requestを自動生成
docker_request_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/request/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./api/presentation/request...

# responseを自動生成
docker_response_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/response/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./api/presentation/response...

# 全てのdomainを自動生成
docker_domain_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/entity/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./domain/entity/...
	docker compose -f docker-compose.local.yml exec gen go generate ./script/repository/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./domain/repository/...
	docker compose -f docker-compose.local.yml exec gen go generate ./script/dao/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./infra/dao/...
	docker compose -f docker-compose.local.yml exec gen go generate ./script/sql/gen.go

# modelを自動生成
docker_entity_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/entity/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./domain/entity/...

# repositoryを自動生成
docker_repository_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/repository/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./domain/repository/...

# daoを自動生成
docker_dao_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/dao/gen.go
	docker compose -f docker-compose.local.yml exec gen go fmt ./infra/dao/...

# sqlを自動生成
docker_sql_gen:
	docker compose -f docker-compose.local.yml exec gen go generate ./script/sql/gen.go

# Swaggerを自動生成
docker_swag_gen:
	docker compose -f docker-compose.local.yml exec api swag init --dir=api --output=docs/swagger/api
	docker compose -f docker-compose.local.yml exec auth swag init --dir=auth --output=docs/swagger/auth

# Swaggerのモックサーバーを起動
docker_swag_mock:
	docker compose -f docker-compose.local.yml exec swagger prism mock ./docs/swagger/api/swagger.yaml --port=8000 --host=0.0.0.0

# ローカルDBに接続
docker_db_user:
	docker compose -f docker-compose.local.yml exec db_user mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user

docker_db_master:
	docker compose -f docker-compose.local.yml exec db_master mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_master

