package item

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

func TestNewItemService_NewItemService(t *testing.T) {
	type args struct {
		userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository
		masterItemMysqlRepository  masterItem.MasterItemMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ItemService
	}{
		{
			name: "正常",
			args: args{
				userItemBoxMysqlRepository: nil,
				masterItemMysqlRepository:  nil,
			},
			want: &itemService{
				userItemBoxMysqlRepository: nil,
				masterItemMysqlRepository:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewItemService(
				tt.args.userItemBoxMysqlRepository,
				tt.args.masterItemMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_Create(t *testing.T) {
	type fields struct {
		userItemBoxMysqlRepository func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository
		masterItemMysqlRepository  func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *ItemCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ItemCreateResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合（既に取得したことがあるアイテム）",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        2,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        2,
							},
							nil,
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want: &ItemCreateResponse{
				UserItemBox: &userItemBox.UserItemBox{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        2,
				},
				MasterItem: &masterItem.MasterItem{
					Id:           1,
					Name:         "テストアイテム",
					ResourceType: enum.ResourceType_Normal,
					RarityType:   enum.RarityType_N,
					Content:      "テストノーマルアイテム",
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：作成できる場合（取得したことがないアイテム）",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want: &ItemCreateResponse{
				UserItemBox: &userItemBox.UserItemBox{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
				MasterItem: &masterItem.MasterItem{
					Id:           1,
					Name:         "テストアイテム",
					ResourceType: enum.ResourceType_Normal,
					RarityType:   enum.RarityType_N,
					Content:      "テストノーマルアイテム",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterItemMysqlRepository.Find",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterItemMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxMysqlRepository.FindOrNil",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userItemBoxMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxMysqlRepository.Update",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        2,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userItemBoxMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxMysqlRepository.Create",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemCreateRequest{
					UserId:       "0:WntR-PyhOJeDiE5jodeR",
					MasterItemId: 1,
					Count:        1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userItemBoxMysqlRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				userItemBoxMysqlRepository: tt.fields.userItemBoxMysqlRepository(ctrl),
				masterItemMysqlRepository:  tt.fields.masterItemMysqlRepository(ctrl),
			}

			got, err := s.Create(tt.args.ctx, tt.args.tx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_Receive(t *testing.T) {
	type fields struct {
		userItemBoxMysqlRepository func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository
		masterItemMysqlRepository  func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *ItemReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ItemReceiveResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合（取得したことがないアイテム））",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 2,
								Count:        2,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 2,
								Count:        2,
							},
							nil,
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム1",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(2),
						).
						Return(
							&masterItem.MasterItem{
								Id:           2,
								Name:         "テストアイテム2",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム2",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemReceiveRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        1,
						},
						{
							MasterItemId: 2,
							Count:        2,
						},
					},
				},
			},
			want: &ItemReceiveResponse{
				UserItemBoxes: userItemBox.UserItemBoxes{
					{
						UserId:       "0:WntR-PyhOJeDiE5jodeR",
						MasterItemId: 1,
						Count:        1,
					},
					{
						UserId:       "0:WntR-PyhOJeDiE5jodeR",
						MasterItemId: 2,
						Count:        2,
					},
				},
				MasterItems: masterItem.MasterItems{
					{
						Id:           1,
						Name:         "テストアイテム1",
						ResourceType: enum.ResourceType_Normal,
						RarityType:   enum.RarityType_N,
						Content:      "テストノーマルアイテム1",
					},
					{
						Id:           2,
						Name:         "テストアイテム2",
						ResourceType: enum.ResourceType_Normal,
						RarityType:   enum.RarityType_N,
						Content:      "テストノーマルアイテム2",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：作成できる場合（既に取得したことがあるアイテム）",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        2,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        2,
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(2),
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 2,
								Count:        2,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 2,
								Count:        4,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 2,
								Count:        4,
							},
							nil,
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム1",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(2),
						).
						Return(
							&masterItem.MasterItem{
								Id:           2,
								Name:         "テストアイテム2",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム2",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemReceiveRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        1,
						},
						{
							MasterItemId: 2,
							Count:        2,
						},
					},
				},
			},
			want: &ItemReceiveResponse{
				UserItemBoxes: userItemBox.UserItemBoxes{
					{
						UserId:       "0:WntR-PyhOJeDiE5jodeR",
						MasterItemId: 1,
						Count:        2,
					},
					{
						UserId:       "0:WntR-PyhOJeDiE5jodeR",
						MasterItemId: 2,
						Count:        4,
					},
				},
				MasterItems: masterItem.MasterItems{
					{
						Id:           1,
						Name:         "テストアイテム1",
						ResourceType: enum.ResourceType_Normal,
						RarityType:   enum.RarityType_N,
						Content:      "テストノーマルアイテム1",
					},
					{
						Id:           2,
						Name:         "テストアイテム2",
						ResourceType: enum.ResourceType_Normal,
						RarityType:   enum.RarityType_N,
						Content:      "テストノーマルアイテム2",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.Create",
			fields: fields{
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
						).
						Return(
							&userItemBox.UserItemBox{
								UserId:       "0:WntR-PyhOJeDiE5jodeR",
								MasterItemId: 1,
								Count:        1,
							},
							nil,
						)
					return m
				},
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								Id:           1,
								Name:         "テストアイテム1",
								ResourceType: enum.ResourceType_Normal,
								RarityType:   enum.RarityType_N,
								Content:      "テストノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemReceiveRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        1,
						},
						{
							MasterItemId: 2,
							Count:        2,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.Create: failed to s.masterItemMysqlRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				userItemBoxMysqlRepository: tt.fields.userItemBoxMysqlRepository(ctrl),
				masterItemMysqlRepository:  tt.fields.masterItemMysqlRepository(ctrl),
			}

			got, err := s.Receive(tt.args.ctx, tt.args.tx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Receive() = %v, want %v", got, tt.want)
			}
		})
	}
}
