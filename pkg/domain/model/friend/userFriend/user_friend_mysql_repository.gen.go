// Package userFriend ユーザーフレンド
//
//go:generate mockgen -source=./user_friend_mysql_repository.gen.go -destination=./user_friend_mysql_repository_mock.gen.go -package=userFriend
package userFriend

import (
	"context"

	"gorm.io/gorm"
)

type UserFriendMysqlRepository interface {
	Find(ctx context.Context, userId string, friendUserId string) (*UserFriend, error)
	FindOrNil(ctx context.Context, userId string, friendUserId string) (*UserFriend, error)
	FindList(ctx context.Context, userId string) (UserFriends, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserFriend) (*UserFriend, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserFriends) (UserFriends, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserFriend) (*UserFriend, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserFriend) error
}
