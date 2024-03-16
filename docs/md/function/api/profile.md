# Profile
プロフィール関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/profile)  

- [Get](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/profile.md#Get)
- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/profile.md#Create)
- [Update](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/profile.md#Update)

## Get
プロフィールを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:ZJJrANH5F8gbNbusyH-9",
}
```
- response
```json
{
    "user_profile": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "プロフィール名",
        "content": "プロフィール内容"
    }
}
```

## Create
プロフィールを作成する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:ZJJrANH5F8gbNbusyH-9",
    "name": "プロフィール名",
    "content": "プロフィール内容"
}
```
- response
```json
{
    "user_profile": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "プロフィール名",
        "content": "プロフィール内容"
    }
}
```

## Update
プロフィールを更新する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:ZJJrANH5F8gbNbusyH-9",
    "name": "プロフィール名",
    "content": "プロフィール内容"
}
```
- response
```json
{
    "user_profile": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "プロフィール名",
        "content": "プロフィール内容"
    }
}
```
