# Friend
フレンド関連。
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/friend)

- [Get](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/friend.md#get)
- [Send](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/friend.md#send)
- [Approve](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/friend.md#approve)
- [Disapprove](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/friend.md#disapprove)
- [Delete](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/friend.md#delete)

## Get
フレンドを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *FriendGetRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *FriendGetResponse | レスポンス |
| err | error | エラー |

## Send
フレンド申請を送る。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| req | *FriendSendRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *FriendSendResponse | レスポンス |
| err | error | エラー |

## Approve
フレンド申請を承認する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| req | *FriendApproveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *FriendApproveResponse | レスポンス |
| err | error | エラー |

## Disapprove
フレンド申請を拒否する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| req | *FriendDisapproveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *FriendDisapproveResponse | レスポンス |
| err | error | エラー |

## Delete
フレンド申請を削除する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| req | *FriendDeleteRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *FriendDeleteResponse | レスポンス |
| err | error | エラー |
