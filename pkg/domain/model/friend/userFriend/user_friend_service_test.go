package userFriend

import (
	"reflect"
	"testing"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

func TestNewUserLoginBonus_CheckReceived(t *testing.T) {
	type fields struct {
		UserFriends UserFriends
	}
	type args struct {
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UserFriends
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				UserFriends: UserFriends{
					{
						UserId:       "0:0000",
						FriendUserId: "1:1111",
						FriendType:   enum.FriendType_Approved,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:2222",
						FriendType:   enum.FriendType_Approved,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:3333",
						FriendType:   enum.FriendType_Applying,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:4444",
						FriendType:   enum.FriendType_Approved,
					},
				},
			},
			want: UserFriends{
				{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_Approved,
				},
				{
					UserId:       "0:0000",
					FriendUserId: "1:2222",
					FriendType:   enum.FriendType_Approved,
				},
				{
					UserId:       "0:0000",
					FriendUserId: "1:4444",
					FriendType:   enum.FriendType_Approved,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserFriends.GetFriends()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
