# kubernetes
k8s_apply:
	make redis_apply
	make mysql_apply
	make api_apply API_NAME=game
	make gen_apply
	make gen_migration
	make gen_master
k8s_delete:
	make api_delete API_NAME=game
	make mysql_delete
	make redis_delete
	make gen_delete

# kubernetes MySQL
mysql_apply:
	kubectl apply -f ./platform/kubernetes/mysql/persistent-volume.yaml
	kubectl apply -f ./platform/kubernetes/mysql/persistent-volume-claim.yaml
	kubectl apply -f ./platform/kubernetes/mysql/deployment.yaml
	kubectl apply -f ./platform/kubernetes/mysql/service.yaml
	kubectl create configmap mysql-init-scripts --from-file=./platform/kubernetes/mysql/init/init-script.sql
mysql_delete:
	kubectl delete -f ./platform/kubernetes/mysql/service.yaml
	kubectl delete -f ./platform/kubernetes/mysql/deployment.yaml
	kubectl delete -f ./platform/kubernetes/mysql/persistent-volume-claim.yaml
	kubectl delete -f ./platform/kubernetes/mysql/persistent-volume.yaml
	kubectl delete configmap mysql-init-scripts

# kubernetes MySQLに接続する
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

# kubernetes Redis
redis_apply:
	kubectl apply -f ./platform/kubernetes/redis/deployment.yaml
	kubectl apply -f ./platform/kubernetes/redis/service.yaml
redis_delete:
	kubectl delete -f ./platform/kubernetes/redis/service.yaml
	kubectl delete -f ./platform/kubernetes/redis/deployment.yaml

# kubernetes API
api_apply:
	docker build -f ./platform/docker/api/$(API_NAME)/Dockerfile --target prod -t localhost:gocrafter-api-$(API_NAME)-local .
	kubectl apply -f ./platform/kubernetes/api/$(API_NAME)/namespace.yaml
	kubectl create configmap gocrafter-env-config --namespace=gocrafter-api-game --from-env-file=.env.local
	kubectl apply -f ./platform/kubernetes/api/$(API_NAME)/deployment.yaml
	kubectl apply -f ./platform/kubernetes/api/$(API_NAME)/service.yaml
# kubernetes API
api_delete:
	kubectl delete -f ./platform/kubernetes/api/$(API_NAME)/service.yaml
	kubectl delete -f ./platform/kubernetes/api/$(API_NAME)/deployment.yaml
	kubectl delete configmap gocrafter-env-config --namespace=gocrafter-api-game
	kubectl delete -f ./platform/kubernetes/api/$(API_NAME)/namespace.yaml

# Gen
gen_apply:
	docker run -it --rm -d --name gocrafter-gen-container --env-file .env.local -v .:/go/src/app localhost:gocrafter-gen-local
gen_delete:
	docker stop gocrafter-gen-container

# kubernetes Genに入る
gen_conn:
	docker exec -it gocrafter-gen-container /bin/sh

# テスト
gen_test:
	docker exec -it gocrafter-gen-container /bin/sh -c "go clean -testcache \
	&& go test -v ./api/game/usecase/... \
	&& go test -v ./pkg/domain/model/... \
	"

# kubernetes マイグレーション
gen_migration:
	docker exec -it gocrafter-gen-container /bin/sh -c "go run ./tools/migration/migration.go"

# kubernetes マスターインポート
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
