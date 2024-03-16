# Friend
フレンド関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/friend)  

- [Get](https://github.com/game-core/gocrafter/blob/main/docs/md/function/friend.md#get)
- [Send](https://github.com/game-core/gocrafter/blob/main/docs/md/function/friend.md#send)
- [Approve](https://github.com/game-core/gocrafter/blob/main/docs/md/function/friend.md#approve)
- [Disapprove](https://github.com/game-core/gocrafter/blob/main/docs/md/function/friend.md#disapprove)
- [Delete](https://github.com/game-core/gocrafter/blob/main/docs/md/function/friend.md#delete)

## Get
フレンドを取得する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO"
}
```
- response
```json
{
    "user_friends": [
        {
            "user_id": "0:BAgor1YTA5DdwE3K1UoO",
            "friend_user_id": "1:BAgor1YTA5DdwE3K1UoO",
            "friend_type": "FriendType_Approved"
        },
        {
            "user_id": "0:BAgor1YTA5DdwE3K1UoO",
            "friend_user_id": "1:BAgor1YTA5DdwE3K1Uo1",
            "friend_type": "FriendType_Approved"
        }
        {
            "user_id": "0:BAgor1YTA5DdwE3K1UoO",
            "friend_user_id": "1:BAgor1YTA5DdwE3K1Uo2",
            "friend_type": "FriendType_NotApproved"
        }
    ]
}
```

## Send
フレンド申請を送る。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "friend_user_id": "0:BAgor1YTA5DfajK1so1"
}
```
- response
```json
{
    "user_friend": {
        "user_id":  "0:BAgor1YTA5DdwE3K1UoO",
        "friend_user_id": "0:BAgor1YTA5DfajK1so1",
        "friend_type": "FriendType_Applying"
    }
}
```

## Approve
フレンド申請を承認する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "friend_user_id": "0:BAgor1YTA5DfajK1so1"
}
```
- response
```json
{
    "user_friend": {
        "user_id":  "0:BAgor1YTA5DdwE3K1UoO",
        "friend_user_id": "0:BAgor1YTA5DfajK1so1",
        "friend_type": "FriendType_Approved"
    }
}
```

## Disapprove
フレンド申請を拒否する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "friend_user_id": "0:BAgor1YTA5DfajK1so1"
}
```
- response
```json
{
    "user_friend": {
        "user_id":  "0:BAgor1YTA5DdwE3K1UoO",
        "friend_user_id": "0:BAgor1YTA5DfajK1so1",
        "friend_type": "FriendType_Disapproved"
    }
}
```

## Delete
フレンド申請を削除する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "friend_user_id": "0:BAgor1YTA5DfajK1so1"
}
```
- response
```json
{
    "user_friend": {
        "user_id":  "0:BAgor1YTA5DdwE3K1UoO",
        "friend_user_id": "0:BAgor1YTA5DfajK1so1",
        "friend_type": "FriendType_NotFriend"
    }
}
```


