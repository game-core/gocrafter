DOCKER_COMPOSE = docker compose -f docker-compose.local.yml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build
	$(DOCKER_COMPOSE_TEST) up -d --build

# wireを自動生成
docker_wire_gen:
	$(DOCKER_COMPOSE) exec gen wire admin/di/wire.go
	$(DOCKER_COMPOSE) exec gen wire api/di/wire.go
	$(DOCKER_COMPOSE) exec gen wire auth/di/wire.go

# apiを自動生成
docker_app_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/admin/request/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/admin/response/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/auth/request/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/auth/response/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/api/request/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/api/response/gen.go
	$(DOCKER_COMPOSE) exec gen wire admin/di/wire.go
	$(DOCKER_COMPOSE) exec gen wire auth/di/wire.go
	$(DOCKER_COMPOSE) exec gen wire api/di/wire.go
	$(DOCKER_COMPOSE) exec gen go fmt ./...

# 全てのdomainを自動生成
docker_domain_gen:
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/admin/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/auth/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/entity/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/admin/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/auth/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/repository/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/admin/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/auth/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/dao/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/enum/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/admin/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/auth/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/config/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/master/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./config/generator/sql/user/gen.go
	$(DOCKER_COMPOSE) exec gen go generate ./domain/...
	$(DOCKER_COMPOSE) exec gen go fmt ./...

# Swaggerを自動生成
docker_swag_gen:
	$(DOCKER_COMPOSE) exec api swag init --dir=api --output=docs/swagger/api
	$(DOCKER_COMPOSE) exec auth swag init --dir=auth --output=docs/swagger/auth

# Swaggerのモックサーバーを起動
docker_swag_mock:
	$(DOCKER_COMPOSE) exec swagger prism mock ./docs/swagger/api/swagger.yaml --port=8000 --host=0.0.0.0

# ローカルDBに接続
docker_db_user_1:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_1

docker_db_user_2:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_2

docker_db_master:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_master

docker_db_config:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_config

docker_db_auth:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_auth

docker_db_admin:
	$(DOCKER_COMPOSE) exec db mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_admin
