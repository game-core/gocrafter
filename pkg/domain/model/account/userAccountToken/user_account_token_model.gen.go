// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

type UserAccountTokens []*UserAccountToken

type UserAccountToken struct {
	UserId string
	Token  string
}

func NewUserAccountToken() *UserAccountToken {
	return &UserAccountToken{}
}

func NewUserAccountTokens() UserAccountTokens {
	return UserAccountTokens{}
}

func SetUserAccountToken(userId string, token string) *UserAccountToken {
	return &UserAccountToken{
		UserId: userId,
		Token:  token,
	}
}
