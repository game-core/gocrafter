# Action
ユーザー行動管理。  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/action)

- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/action.md#GetMaster)
- [Check](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/action.md#Check)
- [Run](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/action.md#Run)

## GetMaster
マスターデータを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ActionGetMasterResponse | レスポンス |
| err | error | エラー |

## Check
アクションが実行済か確認する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| now | time.Time | 現在時刻 |
| req | *ActionCheckRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| err | error | エラー |

## Run
アクションを実行する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| now | time.Time | 現在時刻 |
| req | *ActionRunRequest) | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| err | error | エラー |
