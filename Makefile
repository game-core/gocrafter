DOCKER_COMPOSE = docker compose -f docker-compose.local.yml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build
	$(DOCKER_COMPOSE_TEST) up -d --build

# wireを自動生成
docker_wire_gen:
	$(DOCKER_COMPOSE) exec gen wire api/di/wire.go

# 全てのappを自動生成
docker_app_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/request/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/response/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./...

# requestを自動生成
docker_request_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/request/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./api/presentation/request...

# responseを自動生成
docker_response_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/response/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./api/presentation/response...

# 全てのdomainを自動生成
docker_domain_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./domain/...
	$(DOCKER_COMPOSE) exec gen go fmt ./...

# modelを自動生成
docker_entity_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/user/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./domain/entity/...

# repositoryを自動生成
docker_repository_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/user/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./domain/repository/...

# daoを自動生成
docker_dao_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/user/gen.go
	$(DOCKER_COMPOSE) exec gen go fmt ./infra/dao/...

# sqlを自動生成
docker_sql_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/user/gen.go

# Swaggerを自動生成
docker_swag_gen:
	$(DOCKER_COMPOSE) exec api swag init --dir=api --output=docs/swagger/api

# Swaggerのモックサーバーを起動
docker_swag_mock:
	$(DOCKER_COMPOSE) exec swagger prism mock ./docs/swagger/api/swagger.yaml --port=8000 --host=0.0.0.0

# ローカルDBに接続
docker_db_user:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user

# ローカルDBに接続
docker_db_master:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_master

