# Item
アイテム関連関連。  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/item)

- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/item.md#Create)
- [Receive](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/item.md#Receive)

## Create
アイテムを作成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *ItemCreateRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ItemCreateResponse | レスポンス |
| err | error | エラー |

## Create
アイテムを受け取る。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *ItemReceiveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *ItemReceiveResponse | レスポンス |
| err | error | エラー |
