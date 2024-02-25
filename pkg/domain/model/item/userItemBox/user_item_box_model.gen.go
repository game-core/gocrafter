// Package userItemBox ユーザーアイテムボックス
package userItemBox

type UserItemBoxs []*UserItemBox

type UserItemBox struct {
	UserId       string
	MasterItemId int64
	Count        int32
}

func NewUserItemBox() *UserItemBox {
	return &UserItemBox{}
}

func NewUserItemBoxs() UserItemBoxs {
	return UserItemBoxs{}
}

func SetUserItemBox(userId string, masterItemId int64, count int32) *UserItemBox {
	return &UserItemBox{
		UserId:       userId,
		MasterItemId: masterItemId,
		Count:        count,
	}
}
