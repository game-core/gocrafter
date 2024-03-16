# Profile
プロフィール関連。
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/profile)  

- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/profile.md#create)
- [Update](https://github.com/game-core/gocrafter/blob/main/docs/md/function/profile.md#update)

## Get
プロフィールを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *ProfileGetRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ProfileGetResponse | レスポンス |
| err | error | エラー |

## Create
プロフィールを作成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *ProfileCreateRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ProfileCreateResponse | レスポンス |
| err | error | エラー |

## Update
プロフィールを更新する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *ProfileUpdateRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ProfileUpdateResponse | レスポンス |
| err | error | エラー |
