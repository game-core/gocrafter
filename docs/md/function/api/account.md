# Account
アカウント関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/account)  

- [Create](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/account.md#create)
- [Login](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/account.md#login)
- [Check](https://github.com/game-core/gocrafter/blob/main/docs/md/function/api/account.md#check)

## Create
アカウントを作成する。
- request
```json
{
    "name": "ユーザー1"
}
```
- response
```json
{
    "user_account": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "ユーザー1",
        "password": "n8zm6ZgNg_4uazTbt4CD",
        "login_at": {
            "seconds": "1710568860",
            "nanos": 199408602
        },
        "logout_at": {
            "seconds": "1710568860",
            "nanos": 199408673
        }
    }
}
```

## Login
アカウントをログインする。
- request
```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
    "name": "ユーザー1",
    "password": "n8zm6ZgNg_4uazTbt4CD"
}
```
- response
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA4MjgyODYsIm5hbWUiOiLjg6bjg7zjgrbjg7wxIiwidXNlcklkIjoiMDpCQWdvcjFZVEE1RGR3RTNLMVVvTyJ9.8ZvYDtNQDbQ0egKn1Qx1OArppISj95rzMh3ARxDTDtQ",
    "user_account": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "ユーザー1",
        "password": "$2a$10$hXU3vhvHF04dj1e5N1GhQOJrnPTo95nq74UojoNg1D1mo5xxKs8m.",
        "login_at": {
            "seconds": "1710569086",
            "nanos": 878253584
        },
        "logout_at": {
            "seconds": "1710568860",
            "nanos": 0
        }
    }
}
```

## Check
アカウントを確認する。
- request

| Key | Value | Description |
| :--- | :--- | :--- |
| Authorization | Bearer eyJhbG... | 認証トークン |

```json
{
    "user_id": "0:BAgor1YTA5DdwE3K1UoO",
}
```
- response
```json
{
    "user_account": {
        "user_id": "0:BAgor1YTA5DdwE3K1UoO",
        "name": "ユーザー1",
        "password": "$2a$10$hXU3vhvHF04dj1e5N1GhQOJrnPTo95nq74UojoNg1D1mo5xxKs8m.",
        "login_at": {
            "seconds": "1710569368",
            "nanos": 0
        },
        "logout_at": {
            "seconds": "1710568860",
            "nanos": 0
        }
    }
}
```
