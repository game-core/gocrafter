DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml
DOCKER_COMPOSE_TEST = docker compose -f docker-compose.test.yaml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build
docker_test_up:
	$(DOCKER_COMPOSE_TEST) up -d --build

# テスト
docker_test_run:
	$(DOCKER_COMPOSE_TEST) exec api-game-test go clean -testcache
	$(DOCKER_COMPOSE_TEST) exec api-game-test go test -v ./api/game/usecase/...
	$(DOCKER_COMPOSE_TEST) exec api-game-test go test -v ./pkg/domain/model/...

# DBに接続
docker_db_user0:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_0
docker_db_user1:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_1
docker_db_master:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_master
docker_db_common:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_common

# ジェネレータに接続
docker_gen:
	$(DOCKER_COMPOSE) exec generator bash

# mockを生成
docker_gen_mock:
	$(DOCKER_COMPOSE) exec generator go generate ./pkg/domain/...

# diを生成
docker_gen_di:
	$(DOCKER_COMPOSE) exec generator wire api/game/di/wire.go

# apiを生成
docker_gen_api:
	$(DOCKER_COMPOSE) exec generator go generate ./tools/generator/api/game/main.go
	$(DOCKER_COMPOSE) exec generator sh ./scripts/bulk-generate-protos.sh
	$(DOCKER_COMPOSE) exec generator goimports -w ./api

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

# マスターインポート
docker_master_import:
	$(DOCKER_COMPOSE) exec generator go run ./tools/masterImport/main.go

# fmt
docker_fmt:
	$(DOCKER_COMPOSE) exec generator goimports -w .
