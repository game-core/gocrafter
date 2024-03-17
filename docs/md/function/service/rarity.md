# Rarity
レアリティ関連。

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/rarity)

- [GetAll](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/rarity.md#GetAll)
- [GetByRarityType](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/rarity.md#GetByRarityType)

## GetAll
レアリティ一覧を取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | masterRarity.MasterRarities | レスポンス |
| err | error | エラー |

## GetByRarityType
レアリティを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| rarityType | enum.RarityType | レアリティタイプ |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *masterRarity.MasterRarity | レスポンス |
| err | error | エラー |
