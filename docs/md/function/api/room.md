# Room
ルーム関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/gocrafter-proto/api/game/room)  

- [Search](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/room.md#Search)
- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/room.md#Create)
- [Delete](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/room.md#Delete)
- [CheckIn](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/room.md#CheckIn)
- [CheckOut](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/room.md#Checkout)

## Search
ルームを検索する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:1PL-CGARg5DePVGI9bz2",
    "name": "ルーム1"
}
```
- response
```json
{
    "common_rooms": [
        {
            "room_id": "C4WWpS_9QgFVyqK0ZS3V",
            "host_user_id": "0:BAgor1YTA5DdwE3K1UoO",
            "room_release_type": "Public",
            "name": "ルーム1",
            "user_count": 1
        },
        {
            "room_id": "ehvpbYYPqseNqx_zn7RQ",
            "host_user_id": "0:BAgor1YTA5DdwE3K1UoO",
            "room_release_type": "Public",
            "name": "ルーム1",
            "user_count": 2
        }
    ]
}
```

## Create
ルームを作成する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "name": "ルーム1",
    "room_release_type": "Public"
}
```
- response
```json
{
    "common_room": {
        "room_id": "C4WWpS_9QgFVyqK0ZS3V",
        "host_user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "room_release_type": "Public",
        "name": "ルーム1",
        "user_count": 1
    }
}
```

## Delete
ルームを作成する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO", 
    "room_id": "C4WWpS_9QgFVyqK0ZS3V"
}
```
- response
```json
{
    "common_room": {
        "room_id": "C4WWpS_9QgFVyqK0ZS3V",
        "host_user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "room_release_type": "Public",
        "name": "ルーム1",
        "user_count": 1
    }
}
```

## CheckIn
ルームに入室する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO", 
    "room_id": "C4WWpS_9QgFVyqK0ZS3V"
}
```
- response
```json
{
    "common_room": {
        "room_id": "C4WWpS_9QgFVyqK0ZS3V",
        "host_user_id": "0:BAgor1YTA5DdwE3Kcsdfa",
        "room_release_type": "Public",
        "name": "ルーム1",
        "user_count": 1
    }
}
```

## CheckOut
ルームを退出する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO", 
    "room_id": "C4WWpS_9QgFVyqK0ZS3V"
}
```
- response
```json
{
    "common_room": {
        "room_id": "C4WWpS_9QgFVyqK0ZS3V",
        "host_user_id": "0:BAgor1YTA5DdwE3Kcsdfa",
        "room_release_type": "Public",
        "name": "ルーム1",
        "user_count": 1
    }
}
```
