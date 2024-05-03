# k8s 環境変数
env_apply:
	kubectl create configmap gocrafter-env-config --from-env-file=.env.local
env_delete:
	kubectl delete configmap gocrafter-env-config

# k8s MySQL
mysql_apply:
	kubectl apply -f ./platform/k8s/mysql/persistent-volume.yaml
	kubectl apply -f ./platform/k8s/mysql/persistent-volume-claim.yaml
	kubectl apply -f ./platform/k8s/mysql/deployment.yaml
	kubectl apply -f ./platform/k8s/mysql/service.yaml
	kubectl create configmap mysql-init-scripts --from-file=./platform/k8s/mysql/init/init-script.sql
mysql_delete:
	kubectl delete -f ./platform/k8s/mysql/service.yaml
	kubectl delete -f ./platform/k8s/mysql/deployment.yaml
	kubectl delete -f ./platform/k8s/mysql/persistent-volume-claim.yaml
	kubectl delete -f ./platform/k8s/mysql/persistent-volume.yaml
	kubectl delete configmap mysql-init-scripts

# k8s MySQLに接続する
mysql_conn:
	kubectl exec -it $(shell kubectl get pods -l app=mysql -o=jsonpath='{.items[0].metadata.name}') -- mysql --host=localhost --user=mysql_user --password=mysql_password
mysql_conn_user0:
	kubectl exec -it $(shell kubectl get pods -l app=mysql -o=jsonpath='{.items[0].metadata.name}') -- mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_0
mysql_conn_user1:
	kubectl exec -it $(shell kubectl get pods -l app=mysql -o=jsonpath='{.items[0].metadata.name}') --  mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_user_1
mysql_conn_master:
	kubectl exec -it $(shell kubectl get pods -l app=mysql -o=jsonpath='{.items[0].metadata.name}') -- mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_master
mysql_conn_common:
	kubectl exec -it $(shell kubectl get pods -l app=mysql -o=jsonpath='{.items[0].metadata.name}') -- mysql --host=localhost --user=mysql_user --password=mysql_password gocrafter_common

# k8s Redis
redis_apply:
	kubectl apply -f ./platform/k8s/redis/deployment.yaml
	kubectl apply -f ./platform/k8s/redis/service.yaml
redis_delete:
	kubectl delete -f ./platform/k8s/redis/service.yaml
	kubectl delete -f ./platform/k8s/redis/deployment.yaml

# k8s API
api_apply:
	docker build -f ./platform/docker/api/$(API_NAME)/Dockerfile --target prod -t localhost:gocrafter-api-$(API_NAME)-local .
	kubectl apply -f ./platform/k8s/api/$(API_NAME)/deployment.yaml
	kubectl apply -f ./platform/k8s/api/$(API_NAME)/service.yaml
# k8s API
api_delete:
	kubectl delete -f ./platform/k8s/api/$(API_NAME)/service.yaml
	kubectl delete -f ./platform/k8s/api/$(API_NAME)/deployment.yaml

# Gen
gen_apply:
	docker run -it --rm -d --name gocrafter-gen-container --env-file .env.local -v .:/go/src/app localhost:gocrafter-gen-local
gen_delete:
	docker stop gocrafter-gen-container

# k8s Genに入る
gen_conn:
	docker exec -it gocrafter-gen-container /bin/sh

# テスト
gen_test:
	docker exec -it gocrafter-gen-container /bin/sh -c "go clean -testcache \
	&& go test -v ./api/game/usecase/... \
	&& go test -v ./pkg/domain/model/... \
	"

# k8s マイグレーション
gen_migration:
	docker exec -it gocrafter-gen-container /bin/sh -c "go run ./tools/migration/migration.go"

# k8s マスターインポート
gen_master:
	docker exec -it gocrafter-gen-container /bin/sh -c "go run ./tools/masterImport/main.go"

# fmt
gen_fmt:
	docker exec -it gocrafter-gen-container /bin/sh -c "goimports -w ."

# mockを生成
gen_mock:
	docker exec -it gocrafter-gen-container /bin/sh -c "go generate ./pkg/domain/..."

# diを生成
gen_di:
	docker exec -it gocrafter-gen-container /bin/sh -c "wire api/game/di/wire.go \
	&& wire api/multi/di/wire.go \
	"

# apiを生成
gen_api:
	docker exec -it gocrafter-gen-container /bin/sh -c "go generate ./tools/generator/api/game/main.go \
	&& go generate ./tools/generator/api/multi/main.go \
	&& sh ./scripts/bulk-generate-protos.sh \
	&& goimports -w ./api \
	"

# domainを生成
gen_domain:
	docker exec -it gocrafter-gen-container /bin/sh -c "go generate ./tools/generator/pkg/domain/enum/main.go \
	&& go generate ./tools/generator/pkg/domain/model/main.go \
	&& goimports -w ./pkg/domain \
	"

# infraを生成
gen_infra:
	docker exec -it gocrafter-gen-container /bin/sh -c "go generate ./tools/generator/pkg/infrastructure/mysql/common/main.go \
	&& go generate ./tools/generator/pkg/infrastructure/mysql/master/main.go \
	&& go generate ./tools/generator/pkg/infrastructure/mysql/user/main.go \
	&& go generate ./tools/generator/pkg/infrastructure/redis/common/main.go \
	&& go generate ./tools/generator/pkg/infrastructure/redis/user/main.go \
	&& goimports -w ./pkg/infrastructure \
	&& goimports -w ./pkg/domain \
	"

# sqlを生成
gen_sql:
	docker exec -it gocrafter-gen-container /bin/sh -c "go generate ./tools/generator/sql/main.go"
