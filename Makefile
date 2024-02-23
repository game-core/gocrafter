DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yaml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build

# ジェネレータに接続
docker_gen:
	$(DOCKER_COMPOSE) exec generator bash

# domainを生成
docker_gen_domain:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/domain/model/main.go
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/domain

# infraを生成
docker_gen_infra:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/common/main.go
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/master/main.go
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/infrastructure/mysql/user/main.go
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/infrastructure
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/domain
