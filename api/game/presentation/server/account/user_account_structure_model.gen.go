// Package account ユーザーアカウント
package account

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetUserAccount(userId string, name string, password string, token string, loginAt *timestamppb.Timestamp, logoutAt *timestamppb.Timestamp) *UserAccount {
	return &UserAccount{
		UserId:   userId,
		Name:     name,
		Password: password,
		Token:    token,
		LoginAt:  loginAt,
		LogoutAt: logoutAt,
	}
}
