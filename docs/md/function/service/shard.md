# Shard
DBシャード関連  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/shard)

- [GetShardKey](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/shard.md#GetShardKey)

## GetShardKey
シャードキーを取得して更新する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | string | レスポンス |
| err | error | エラー |
