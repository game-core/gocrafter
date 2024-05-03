# 環境構築

## Server

- Kubernetesを起動
```
make k8s_apply
```
- マイグレーション
  - `./docs/sql`配下のsqlファイルが実行される
```
make gen_migration
```
- マスターデータインポート
  - 事前に[GASの設定]()を行う
```
make dgen_master
```

## Test
- テスト実行
```
make gen_test
```
