# Friend
フレンド関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/friend)  

## Get
フレンドを取得する。
- request
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
