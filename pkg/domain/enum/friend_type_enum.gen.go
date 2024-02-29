// Package enum フレンドタイプ
package enum

type FriendType int32

const (
	FriendType_Applying    FriendType = 0
	FriendType_NotApproved FriendType = 1
	FriendType_Approved    FriendType = 2
	FriendType_Disapproved FriendType = 3
	FriendType_NotFriend   FriendType = 4
)
