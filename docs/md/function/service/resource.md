# Resource
リソース関連。

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/resource)

- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/resource.md#GetAll)
- [Receive](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/resource.md#GetByResourceType)

## GetAll
リソース一覧を取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | masterResource.MasterResources | レスポンス |
| err | error | エラー |

## GetByResourceType
リソースを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| resourceType | enum.ResourceType | リソースタイプ |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *masterResource.MasterResource | レスポンス |
| err | error | エラー |
