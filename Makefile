DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yaml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build

# ジェネレータに接続
docker_gen:
	$(DOCKER_COMPOSE) exec generator bash

# mockを生成
docker_gen_mock:
	$(DOCKER_COMPOSE) exec generator go generate ./pkg/domain/...

# apiを生成
docker_gen_api:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/api/game/main.go
	$(DOCKER_COMPOSE) exec generator sh ./scripts/bulk-generate-protos.sh

# domainを生成
docker_gen_domain:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/domain/enum/main.go
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/domain/model/main.go
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/domain

# infraを生成
docker_gen_infra:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/common/main.go
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/master/main.go
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/user/main.go
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/infrastructure
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/domain

# sqlを生成
docker_gen_sql:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/sql/main.go

# マイグレーション
docker_migration:
	$(DOCKER_COMPOSE) exec generator go run ./tools/migration/migration.go

# fmt
docker_fmt:
	$(DOCKER_COMPOSE) exec generator goimports -w .
