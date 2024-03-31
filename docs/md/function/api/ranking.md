# Ranking
ランキング関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/ranking)  

- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/ranking.md#GetMaster)
- [Get](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/ranking.md#Get)
- [Update](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/ranking.md#Update)

## GetMaster
ランキングのマスターデータを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
  "user_id": "0:1PL-CGARg5DePVGI9bz2",
  "master_ranking_id": "1"
}
```
- response
```json
{
  "master_ranking": {
    "id": "1",
    "master_ranking_event_id": "1",
    "name": "ワールドランキング",
    "ranking_scope_type": "Room",
    "ranking_limit": 3
  },
  "master_ranking_event": {
    "id": "1",
    "name": "ランキングイベント",
    "reset_hour": 9,
    "interval_hour": 24,
    "repeat_setting": true,
    "start_at": {
      "seconds": "1707004800",
      "nanos": 0
    }
  },
  "master_ranking_scope": {
    "id": "1",
    "name": "ルームランキング",
    "ranking_scope_type": "Room"
  }
}
```

## Get
ランキングを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
  "user_id": "0:1PL-CGARg5DePVGI9bz2",
  "master_ranking_id": "1",
  "room_id": "ehvpbYYPqseNqx_zn7RQ"
}
```
- response
```json
{
  "common_ranking_rooms": [
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz0",
      "score": 128,
      "ranked_at": {
        "seconds": "1711873561",
        "nanos": 0
      }
    },
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz2",
      "score": 128,
      "ranked_at": {
        "seconds": "1711876125",
        "nanos": 0
      }
    },
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz1",
      "score": 127,
      "ranked_at": {
        "seconds": "1711872824",
        "nanos": 0
      }
    }
  ],
  "common_ranking_worlds": []
}
```

## Update
ランキングを更新する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
  "user_id": "0:1PL-CGARg5DePVGI9bz2",
  "master_ranking_id": "1",
  "score": 129,
  "room_id": "ehvpbYYPqseNqx_zn7RQ"
}
```
- response
```json
{
  "common_ranking_rooms": [
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz0",
      "score": 128,
      "ranked_at": {
        "seconds": "1711873561",
        "nanos": 0
      }
    },
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz2",
      "score": 129,
      "ranked_at": {
        "seconds": "1711885582",
        "nanos": 480676859
      }
    },
    {
      "master_ranking_id": "1",
      "room_id": "ehvpbYYPqseNqx_zn7RQ",
      "user_id": "0:1PL-CGARg5DePVGI9bz1",
      "score": 127,
      "ranked_at": {
        "seconds": "1711872824",
        "nanos": 0
      }
    }
  ],
  "common_ranking_worlds": []
}
```
