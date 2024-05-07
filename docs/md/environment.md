# 環境構築

## k8sの環境設定
- ドキュメントを参考に構築(kubectlが使えるように慣ればOK)  
[https://kubernetes.io/ja/docs/reference/kubectl/](https://kubernetes.io/ja/docs/home/)


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
