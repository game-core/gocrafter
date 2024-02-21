DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yaml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build

# ジェネレータに接続
docker_gen:
	$(DOCKER_COMPOSE) exec generator bash

# modelを生成
docker_gen_model:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/pkg/domain/model.go
	$(DOCKER_COMPOSE) exec generator goimports -w ./pkg/domain
