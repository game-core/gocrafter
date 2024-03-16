# LoginBonus
ログインボーナス関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/loginBonus)  

- [GetUser](https://github.com/game-core/gocrafter/blob/main/docs/md/function/loginBonus.md#GetUser)
- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/loginBonus.md#GetMaster)
- [Receive](https://github.com/game-core/gocrafter/blob/main/docs/md/function/loginBonus.md#Receive)

## GetUser
ログインボーナスのユーザーステータスを取得する。
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
    "user_login_bonuses": [
        {
            "user_id": "0:ZJJrANH5F8gbNbusyH-9",
            "master_login_bonus_id": "1",
            "received_at": {
                "seconds": "1710562227",
                "nanos": 0
            }
        }
    ]
}
```

## GetMaster
ログインボーナスのマスターデータを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "master_login_bonus_id": 1
}
```
- response
```json
{
    "master_login_bonus_items": [
        {
            "id": "1",
            "master_login_bonus_schedule_id": "1",
            "master_item_id": "1",
            "name": "ログイン1日目のアイテム",
            "count": 1
        },
        {
            "id": "2",
            "master_login_bonus_schedule_id": "2",
            "master_item_id": "1",
            "name": "ログイン2日目のアイテム",
            "count": 1
        },
        {
            "id": "3",
            "master_login_bonus_schedule_id": "3",
            "master_item_id": "1",
            "name": "ログイン3日目のアイテム",
            "count": 1
        },
        {
            "id": "4",
            "master_login_bonus_schedule_id": "4",
            "master_item_id": "1",
            "name": "ログイン4日目のアイテム",
            "count": 1
        },
        {
            "id": "5",
            "master_login_bonus_schedule_id": "5",
            "master_item_id": "1",
            "name": "ログイン5日目のアイテム",
            "count": 1
        },
        {
            "id": "6",
            "master_login_bonus_schedule_id": "6",
            "master_item_id": "1",
            "name": "ログイン6日目のアイテム",
            "count": 1
        },
        {
            "id": "7",
            "master_login_bonus_schedule_id": "7",
            "master_item_id": "1",
            "name": "ログイン7日目のアイテム",
            "count": 1
        },
        {
            "id": "8",
            "master_login_bonus_schedule_id": "7",
            "master_item_id": "2",
            "name": "ログイン7日目の追加アイテム",
            "count": 2
        }
    ],
    "master_login_bonus_schedules": [
        {
            "id": "1",
            "master_login_bonus_id": "1",
            "step": 0,
            "name": "ログイン1日目"
        },
        {
            "id": "2",
            "master_login_bonus_id": "1",
            "step": 1,
            "name": "ログイン2日目"
        },
        {
            "id": "3",
            "master_login_bonus_id": "1",
            "step": 2,
            "name": "ログイン3日目"
        },
        {
            "id": "4",
            "master_login_bonus_id": "1",
            "step": 3,
            "name": "ログイン4日目"
        },
        {
            "id": "5",
            "master_login_bonus_id": "1",
            "step": 4,
            "name": "ログイン5日目"
        },
        {
            "id": "6",
            "master_login_bonus_id": "1",
            "step": 5,
            "name": "ログイン6日目"
        },
        {
            "id": "7",
            "master_login_bonus_id": "1",
            "step": 6,
            "name": "ログイン7日目"
        }
    ],
    "master_login_bonus": {
        "id": "1",
        "master_login_bonus_event_id": "1",
        "name": "通常ログインボーナス"
    },
    "master_login_bonus_event": {
        "id": "1",
        "name": "通常ログインボーナスイベント",
        "reset_hour": 9,
        "interval_hour": 24,
        "repeat_setting": true,
        "start_at": {
            "seconds": "1707004800",
            "nanos": 0
        }
    }
}
```

## Receive
ログインボーナスを受けとる。
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
    "master_login_bonus_items": [
        {
            "id": "7",
            "master_login_bonus_schedule_id": "7",
            "master_item_id": "1",
            "name": "ログイン7日目のアイテム",
            "count": 1
        },
        {
            "id": "8",
            "master_login_bonus_schedule_id": "7",
            "master_item_id": "2",
            "name": "ログイン7日目の追加アイテム",
            "count": 2
        }
    ],
    "user_login_bonus": {
        "user_id": "0:ZJJrANH5F8gbNbusyH-9",
        "master_login_bonus_id": "1",
        "received_at": {
            "seconds": "1710605245",
            "nanos": 587063179
        }
    },
    "master_login_bonus": {
        "id": "1",
        "master_login_bonus_event_id": "1",
        "name": "通常ログインボーナス"
    },
    "master_login_bonus_event": {
        "id": "1",
        "name": "通常ログインボーナスイベント",
        "reset_hour": 9,
        "interval_hour": 24,
        "repeat_setting": true,
        "start_at": {
            "seconds": "1707004800",
            "nanos": 0
        }
    },
    "master_login_bonus_schedule": {
        "id": "7",
        "master_login_bonus_id": "1",
        "step": 6,
        "name": "ログイン7日目"
    }
}
```
