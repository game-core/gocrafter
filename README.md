# gocrafter
ゲームサーバーのテンプレート

## 依存関係
![image](https://github.com/game-core/gocrafter/assets/71867595/fe269513-32aa-45ad-96d4-ff882e103777)

## 環境構築
- コンテナを起動
```
make docker_up
```
- DBコンテナに接続（DBを更新する場合）
```
make docker_db_master
make_docker_db_user_1
make_docker_db_user_2
make_docker_db_config
make_docker_db_admin
```

## 自動生成
- domain
```
make docker_domain_gen
```
- request, response
```
make docker_app_gen
```
- wire
```
make docker_wire_gen
```
