# IdleBonus
放置ボーナス関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/idleBonus)  

- [GetUser](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/idleBonus.md#GetUser)
- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/idleBonus.md#GetMaster)
- [Receive](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/idleBonus.md#Receive)

## GetUser
放置ボーナスのユーザーステータスを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "user_id": "0:aP9UvDkOqvP5iW4YRSd6"
}
```
- response
```json
{
    "user_idle_bonuses": [
        {
            "user_id": "0:aP9UvDkOqvP5iW4YRSd6",
            "master_idle_bonus_id": "1",
            "received_at": {
                "seconds": "1710604460",
                "nanos": 0
            }
        }
    ]
}
```

## GetMaster
放置ボーナスのマスターデータを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "user_id": "0:1PL-CGARg5DePVGI9bz2",
    "master_idle_bonus_id": 1
}
```
- response
```json
{
    "master_idle_bonus_items": [
        {
            "id": "1",
            "master_idle_bonus_schedule_id": "1",
            "master_item_id": "1",
            "name": "放置1回目のアイテム",
            "count": 1
        },
        {
            "id": "2",
            "master_idle_bonus_schedule_id": "2",
            "master_item_id": "1",
            "name": "放置2回目のアイテム",
            "count": 1
        },
        {
            "id": "3",
            "master_idle_bonus_schedule_id": "3",
            "master_item_id": "1",
            "name": "放置3回目のアイテム",
            "count": 1
        },
        {
            "id": "4",
            "master_idle_bonus_schedule_id": "4",
            "master_item_id": "1",
            "name": "放置4回目のアイテム",
            "count": 1
        },
        {
            "id": "5",
            "master_idle_bonus_schedule_id": "5",
            "master_item_id": "1",
            "name": "放置5回目のアイテム",
            "count": 1
        },
        {
            "id": "6",
            "master_idle_bonus_schedule_id": "6",
            "master_item_id": "1",
            "name": "放置6回目のアイテム",
            "count": 1
        },
        {
            "id": "7",
            "master_idle_bonus_schedule_id": "7",
            "master_item_id": "1",
            "name": "放置7回目のアイテム",
            "count": 1
        },
        {
            "id": "8",
            "master_idle_bonus_schedule_id": "7",
            "master_item_id": "2",
            "name": "放置7回目の追加アイテム",
            "count": 2
        }
    ],
    "master_idle_bonus_schedules": [
        {
            "id": "1",
            "master_idle_bonus_id": "1",
            "step": 0,
            "name": "放置1回目"
        },
        {
            "id": "2",
            "master_idle_bonus_id": "1",
            "step": 1,
            "name": "放置2回目"
        },
        {
            "id": "3",
            "master_idle_bonus_id": "1",
            "step": 2,
            "name": "放置3回目"
        },
        {
            "id": "4",
            "master_idle_bonus_id": "1",
            "step": 3,
            "name": "放置4回目"
        },
        {
            "id": "5",
            "master_idle_bonus_id": "1",
            "step": 4,
            "name": "放置5回目"
        },
        {
            "id": "6",
            "master_idle_bonus_id": "1",
            "step": 5,
            "name": "放置6回目"
        },
        {
            "id": "7",
            "master_idle_bonus_id": "1",
            "step": 6,
            "name": "放置7回目"
        }
    ],
    "master_idle_bonus": {
        "id": "1",
        "master_idle_bonus_event_id": "1",
        "name": "通常放置ボーナス"
    },
    "master_idle_bonus_event": {
        "id": "1",
        "name": "通常放置ボーナスイベント",
        "reset_hour": 9,
        "interval_hour": 1,
        "repeat_setting": true,
        "start_at": {
            "seconds": "1707004800",
            "nanos": 0
        }
    }
}
```

## Receive
放置ボーナスを受けとる。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "user_id": "0:aP9UvDkOqvP5iW4YRSd6",
    "master_idle_bonus_id": 1
}
```
- response
```json
{
    "master_idle_bonus_items": [
        {
            "id": "1",
            "master_idle_bonus_schedule_id": "1",
            "master_item_id": "1",
            "name": "放置1回目のアイテム",
            "count": 1
        },
        {
            "id": "2",
            "master_idle_bonus_schedule_id": "2",
            "master_item_id": "1",
            "name": "放置2回目のアイテム",
            "count": 1
        },
        {
            "id": "3",
            "master_idle_bonus_schedule_id": "3",
            "master_item_id": "1",
            "name": "放置3回目のアイテム",
            "count": 1
        },
        {
            "id": "4",
            "master_idle_bonus_schedule_id": "4",
            "master_item_id": "1",
            "name": "放置4回目のアイテム",
            "count": 1
        },
        {
            "id": "5",
            "master_idle_bonus_schedule_id": "5",
            "master_item_id": "1",
            "name": "放置5回目のアイテム",
            "count": 1
        },
        {
            "id": "6",
            "master_idle_bonus_schedule_id": "6",
            "master_item_id": "1",
            "name": "放置6回目のアイテム",
            "count": 1
        },
        {
            "id": "7",
            "master_idle_bonus_schedule_id": "7",
            "master_item_id": "1",
            "name": "放置7回目のアイテム",
            "count": 1
        },
        {
            "id": "8",
            "master_idle_bonus_schedule_id": "7",
            "master_item_id": "2",
            "name": "放置7回目の追加アイテム",
            "count": 2
        }
    ],
    "master_idle_bonus_schedules": [
        {
            "id": "1",
            "master_idle_bonus_id": "1",
            "step": 0,
            "name": "放置1回目"
        },
        {
            "id": "2",
            "master_idle_bonus_id": "1",
            "step": 1,
            "name": "放置2回目"
        },
        {
            "id": "3",
            "master_idle_bonus_id": "1",
            "step": 2,
            "name": "放置3回目"
        },
        {
            "id": "4",
            "master_idle_bonus_id": "1",
            "step": 3,
            "name": "放置4回目"
        },
        {
            "id": "5",
            "master_idle_bonus_id": "1",
            "step": 4,
            "name": "放置5回目"
        },
        {
            "id": "6",
            "master_idle_bonus_id": "1",
            "step": 5,
            "name": "放置6回目"
        },
        {
            "id": "7",
            "master_idle_bonus_id": "1",
            "step": 6,
            "name": "放置7回目"
        }
    ],
    "user_idle_bonus": {
        "user_id": "0:aP9UvDkOqvP5iW4YRSd6",
        "master_idle_bonus_id": "1",
        "received_at": {
            "seconds": "1710604613",
            "nanos": 384528768
        }
    },
    "master_idle_bonus": {
        "id": "1",
        "master_idle_bonus_event_id": "1",
        "name": "通常放置ボーナス"
    },
    "master_idle_bonus_event": {
        "id": "1",
        "name": "通常放置ボーナスイベント",
        "reset_hour": 9,
        "interval_hour": 1,
        "repeat_setting": true,
        "start_at": {
            "seconds": "1707004800",
            "nanos": 0
        }
    }
}
```
