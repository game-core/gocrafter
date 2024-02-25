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
		userItemBoxRepository userItemBox.UserItemBoxRepository
		masterItemRepository  masterItem.MasterItemRepository
	}
	tests := []struct {
		name string
		args args
		want ItemService
	}{
		{
			name: "正常",
			args: args{
				userItemBoxRepository: nil,
				masterItemRepository:  nil,
			},
			want: &itemService{
				userItemBoxRepository: nil,
				masterItemRepository:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewItemService(
				tt.args.userItemBoxRepository,
				tt.args.masterItemRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_Create(t *testing.T) {
	type fields struct {
		userItemBoxRepository func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository
		masterItemRepository  func(ctrl *gomock.Controller) masterItem.MasterItemRepository
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
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			name: "異常：s.masterItemRepository.Find",
			fields: fields{
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
					return m
				},
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterItemRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxRepository.FindOrNil",
			fields: fields{
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userItemBoxRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxRepository.Update",
			fields: fields{
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userItemBoxRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userItemBoxRepository.Create",
			fields: fields{
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userItemBoxRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				userItemBoxRepository: tt.fields.userItemBoxRepository(ctrl),
				masterItemRepository:  tt.fields.masterItemRepository(ctrl),
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
		userItemBoxRepository func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository
		masterItemRepository  func(ctrl *gomock.Controller) masterItem.MasterItemRepository
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
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
				userItemBoxRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxRepository {
					m := userItemBox.NewMockUserItemBoxRepository(ctrl)
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
				masterItemRepository: func(ctrl *gomock.Controller) masterItem.MasterItemRepository {
					m := masterItem.NewMockMasterItemRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.Create: failed to s.masterItemRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				userItemBoxRepository: tt.fields.userItemBoxRepository(ctrl),
				masterItemRepository:  tt.fields.masterItemRepository(ctrl),
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
