# Config
設定関連。

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/config)

- [GetAll](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/config.md#GetAll)
- [GetByConfigType](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/config.md#GetByConfigType)

## GetAll
設定一覧を取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | masterConfig.MasterConfigs | レスポンス |
| err | error | エラー |

## GetByConfigType
設定を取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| configType | enum.ConfigType | 設定タイプ |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *masterConfig.MasterConfig | レスポンス |
| err | error | エラー |
