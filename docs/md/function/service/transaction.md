# Transaction
DBトランザクション関連。  

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/transaction)

- [CommonBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonBegin)
- [CommonEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonEnd)
- [MasterBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterBegin)
- [MasterEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterEnd)
- [UserBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserBegin)
- [UserEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserEnd)
- [MultiUserBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserBegin)
- [MultiUserEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserEnd)

## CommonBegin
Common DBのトランザクションを開始する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *gorm.DB | レスポンス |
| err | error | エラー |

## CommonEnd
Common DBのトランザクションを終了する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## MasterBegin
Master DBのトランザクションを開始する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *gorm.DB | レスポンス |
| err | error | エラー |

## MasterEnd
Master DBのトランザクションを終了する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## UserBegin
User DBのトランザクションを開始する。（単一シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| userId | string | ユーザーID |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *gorm.DB | レスポンス |
| err | error | エラー |

## UserEnd
User DBのトランザクションを終了する。（単一シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## MultiUserBegin
User DBのトランザクションを開始する。（複数シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| userIds | []string | ユーザーID一覧 |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | map[string]*gorm.DB | レスポンス |
| err | error | エラー |

## MultiUserEnd
User DBのトランザクションを終了する。（複数シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| err | error | エラー |
