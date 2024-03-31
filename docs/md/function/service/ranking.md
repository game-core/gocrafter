# Ranking

ランキング関連。  
[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/ranking)

- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/ranking.md#GetMaster)
- [Get](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/ranking.md#Get)
- [Update](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/ranking.md#Update)

## GetMaster

ランキングのマスターデータを取得する。

- request

| Name | Type                     | Description |
|:-----|:-------------------------|:------------|
| ctx  | context.Context          | コンテキスト      |
| req  | *RankingGetMasterRequest | リクエスト       |

- response

| Name | Type                      | Description |
|:-----|:--------------------------|:------------|
| res  | *RankingGetMasterResponse | レスポンス       |
| err  | error                     | エラー         |

## Get

ランキングを取得する。

- request

| Name | Type               | Description |
|:-----|:-------------------|:------------|
| ctx  | context.Context    | コンテキスト      |
| now  | time.Time          | 現在時刻        |
| req  | *RankingGetRequest | リクエスト       |

- response

| Name | Type                | Description |
|:-----|:--------------------|:------------|
| res  | *RankingGetResponse | レスポンス       |
| err  | error               | エラー         |

## Update

ランキングを更新する。

- request

| Name | Type                  | Description |
|:-----|:----------------------|:------------|
| ctx  | context.Context       | コンテキスト      |
| tx   | *gorm.DB              | トランザクション    |
| now  | time.Time             | 現在時刻        |
| req  | *RankingUpdateRequest | リクエスト       |

- response

| Name | Type                   | Description |
|:-----|:-----------------------|:------------|
| res  | *RankingUpdateResponse | レスポンス       |
| err  | error                  | エラー         |
