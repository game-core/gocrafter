# 環境構築

## Server

- コンテナ起動
```
make docker_up
```
- マイグレーション
  - `./docs/sql`配下のsqlファイルが実行される
```
make docker_migration
```
- マスターデーターインポート
  - 事前に[GASの設定]()を行う
```
make docker_master_import
```

## Test
- コンテナ起動
```
make docker_test_up
```
- テスト実行
```
make docker_test_run
```
