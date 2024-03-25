# Room

ルーム関連。  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/room)

- [Search](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/room.md#Search)
- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/room.md#Create)
- [Delete](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/room.md#Delete)
- [CheckIn](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/room.md#CheckIn)
- [CheckOut](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/room.md#Checkout)

## Search

ルームを検索する。

- request

| Name | Type              | Description |
|:-----|:------------------|:------------|
| ctx  | context.Context   | コンテキスト      |
| req  | RoomSearchRequest | リクエスト       |

- response

| Name | Type               | Description |
|:-----|:-------------------|:------------|
| res  | RoomSearchResponse | レスポンス       |
| err  | error              | エラー         |

## Create

ルームを作成する。

- request

| Name | Type              | Description |
|:-----|:------------------|:------------|
| ctx  | context.Context   | コンテキスト      |
| tx   | *gorm.DB          | トランザクション    |
| req  | RoomCreateRequest | リクエスト       |

- response

| Name | Type               | Description |
|:-----|:-------------------|:------------|
| res  | RoomCreateResponse | レスポンス       |
| err  | error              | エラー         |

## Delete

ルームを削除する。

- request

| Name | Type              | Description |
|:-----|:------------------|:------------|
| ctx  | context.Context   | コンテキスト      |
| tx   | *gorm.DB          | トランザクション    |
| req  | RoomDeleteRequest | リクエスト       |

- response

| Name | Type               | Description |
|:-----|:-------------------|:------------|
| res  | RoomDeleteResponse | レスポンス       |
| err  | error              | エラー         |

## CheckIn

ルームに入室する。

- request

| Name | Type               | Description |
|:-----|:-------------------|:------------|
| ctx  | context.Context    | コンテキスト      |
| tx   | *gorm.DB           | トランザクション    |
| req  | RoomCheckInRequest | リクエスト       |

- response

| Name | Type                | Description |
|:-----|:--------------------|:------------|
| res  | RoomCheckInResponse | レスポンス       |
| err  | error               | エラー         |

## CheckOut

ルームを退出する。

- request

| Name | Type                | Description |
|:-----|:--------------------|:------------|
| ctx  | context.Context     | コンテキスト      |
| tx   | *gorm.DB            | トランザクション    |
| req  | RoomCheckOutRequest | リクエスト       |

- response

| Name | Type                 | Description |
|:-----|:---------------------|:------------|
| res  | RoomCheckOutResponse | レスポンス       |
| err  | error                | エラー         |
