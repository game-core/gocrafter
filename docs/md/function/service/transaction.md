# Transaction
DBトランザクション関連。  

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/transaction)

- [CommonMysqlBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonMysqlBegin)
- [CommonMysqlEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonMysqlEnd)
- [MasterMysqlBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterMysqlBegin)
- [MasterMysqlEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterMysqlEnd)
- [UserMysqlBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserMysqlBegin)
- [UserMysqlEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserMysqlEnd)
- [MultiUserMysqlBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserMysqlBegin)
- [MultiUserMysqlEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserMysqlEnd)

## CommonMysqlBegin
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

## CommonMysqlEnd
Common DBのトランザクションを終了する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## MasterMysqlBegin
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

## MasterMysqlEnd
Master DBのトランザクションを終了する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## UserMysqlBegin
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

## UserMysqlEnd
User DBのトランザクションを終了する。（単一シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| err | error | エラー |

## MultiUserMysqlBegin
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

## MultiUserMysqlEnd
User DBのトランザクションを終了する。（複数シャード）
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| txs | map[string]*gorm.DB | トランザクションマップ |
| err | error | エラー |
