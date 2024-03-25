# Account
アカウント関連。  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/account)

- [FindByUserId](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/account.md#FindByUserId)
- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/account.md#Create)
- [Login](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/account.md#Login)
- [Check](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/account.md#Check)
- [GenerateUserId](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/account.md#GenerateUserId)

## FindByUserId
ユーザーIDから取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| userId | string | ユーザーID |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *userAccount.UserAccount | ユーザーアカウントモデル |
| err | error | エラー |


## Create
アカウントを作成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *AccountCreateRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *AccountCreateResponse | レスポンス |
| err | error | エラー |

## Login
アカウントをログインする。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| req | *AccountLoginRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *AccountLoginResponse | レスポンス |
| err | error | エラー |

## Check
アカウントを確認する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *AccountCheckRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *AccountCheckResponse | レスポンス |
| err | error | エラー |

## GenerateUserId
ユーザーIDを生成する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| userId | string | ユーザーID |
| err | error | エラー |
