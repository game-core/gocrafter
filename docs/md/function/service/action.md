# Action
ユーザー行動管理。

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
