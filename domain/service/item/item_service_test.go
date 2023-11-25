package item

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/api/presentation/request/item"
	response "github.com/game-core/gocrafter/api/presentation/response/item"
	masterItemEntity "github.com/game-core/gocrafter/domain/entity/master/item"
	userItemEntity "github.com/game-core/gocrafter/domain/entity/user/item"
	masterItemRepository "github.com/game-core/gocrafter/domain/repository/master/item"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	userItemRepository "github.com/game-core/gocrafter/domain/repository/user/item"
)

func TestExampleService_GetItemToEntity(t *testing.T) {
	type fields struct {
		itemRepository func(ctrl *gomock.Controller) masterItemRepository.ItemRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *masterItemEntity.Item
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item",
						).
						Return(
							&masterItemEntity.Item{
								ID:     1,
								Name:   "item",
								Detail: "detail",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				name: "item",
			},
			want: &masterItemEntity.Item{
				ID:     1,
				Name:   "item",
				Detail: "detail",
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（itemRepository.FindByName）",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item",
						).
						Return(
							nil,
							errors.New("itemRepository.FindByName"),
						)
					return m
				},
			},
			args: args{
				name: "item",
			},
			want:    nil,
			wantErr: errors.New("itemRepository.FindByName"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				itemRepository: tt.fields.itemRepository(ctrl),
			}

			got, err := s.GetItemToEntity(tt.args.name)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetItemToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleService_ReceiveItemInBox(t *testing.T) {
	type fields struct {
		transactionRepository func(ctrl *gomock.Controller) userRepository.TransactionRepository
		itemRepository        func(ctrl *gomock.Controller) masterItemRepository.ItemRepository
		itemBoxRepository     func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository
	}
	type args struct {
		req *request.ReceiveItemInBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ReceiveItemInBox
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin("SHARD_1").
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Commit(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							&masterItemEntity.Item{
								ID:        1,
								Name:      "item1",
								Detail:    "detail1",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindByName(
							"item2",
						).
						Return(
							&masterItemEntity.Item{
								ID:        2,
								Name:      "item2",
								Detail:    "detail2",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item1",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item2",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     3,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.ReceiveItemInBox{
					ShardKey:  "SHARD_1",
					AccountID: 1,
					Items: request.Items{
						{
							Name:  "item1",
							Count: 1,
						},
						{
							Name:  "item2",
							Count: 2,
						},
					},
				},
			},
			want: &response.ReceiveItemInBox{
				Status: 200,
				Items: response.Items{
					{
						ID:     1,
						Name:   "item1",
						Detail: "detail1",
						Count:  1,
					},
					{
						ID:     2,
						Name:   "item2",
						Detail: "detail2",
						Count:  2,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：トランザクションエラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin("SHARD_1").
						Return(
							nil,
							errors.New("transactionRepository.Begin"),
						)
					return m
				},
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					return m
				},
			},
			args: args{
				req: &request.ReceiveItemInBox{
					ShardKey:  "SHARD_1",
					AccountID: 1,
					Items: request.Items{
						{
							Name:  "item1",
							Count: 1,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.New("transactionRepository.Begin"),
		},
		{
			name: "異常：エラー（receiveItemInBox）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin("SHARD_1").
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							nil,
							errors.New("receiveItemInBox"),
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					return m
				},
			},
			args: args{
				req: &request.ReceiveItemInBox{
					ShardKey:  "SHARD_1",
					AccountID: 1,
					Items: request.Items{
						{
							Name:  "item1",
							Count: 1,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.New("receiveItemInBox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				transactionRepository: tt.fields.transactionRepository(ctrl),
				itemRepository:        tt.fields.itemRepository(ctrl),
				itemBoxRepository:     tt.fields.itemBoxRepository(ctrl),
			}

			got, err := s.ReceiveItemInBox(tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ReceiveItemInBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceiveItemInBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleService_receiveItemInBox(t *testing.T) {
	type fields struct {
		itemRepository    func(ctrl *gomock.Controller) masterItemRepository.ItemRepository
		itemBoxRepository func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository
	}
	type args struct {
		items     *request.Items
		accountID int64
		shardKey  string
		tx        *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.Items
		wantErr error
	}{
		{
			name: "正常：受け取りできる（初回）",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							&masterItemEntity.Item{
								ID:        1,
								Name:      "item1",
								Detail:    "detail1",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindByName(
							"item2",
						).
						Return(
							&masterItemEntity.Item{
								ID:        2,
								Name:      "item2",
								Detail:    "detail2",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item1",
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item2",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     3,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				items: &request.Items{
					{
						Name:  "item1",
						Count: 1,
					},
					{
						Name:  "item2",
						Count: 2,
					},
				},
				accountID: int64(1),
				shardKey:  "SHARD_1",
				tx:        nil,
			},
			want: &response.Items{
				{
					ID:     1,
					Name:   "item1",
					Detail: "detail1",
					Count:  1,
				},
				{
					ID:     2,
					Name:   "item2",
					Detail: "detail2",
					Count:  2,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：受け取りできる（2回目以降）",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							&masterItemEntity.Item{
								ID:        1,
								Name:      "item1",
								Detail:    "detail1",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindByName(
							"item2",
						).
						Return(
							&masterItemEntity.Item{
								ID:        2,
								Name:      "item2",
								Detail:    "detail2",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item1",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item2",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        2,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item2",
								Count:     3,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				items: &request.Items{
					{
						Name:  "item1",
						Count: 1,
					},
					{
						Name:  "item2",
						Count: 2,
					},
				},
				accountID: int64(1),
				shardKey:  "SHARD_1",
				tx:        nil,
			},
			want: &response.Items{
				{
					ID:     1,
					Name:   "item1",
					Detail: "detail1",
					Count:  1,
				},
				{
					ID:     2,
					Name:   "item2",
					Detail: "detail2",
					Count:  2,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（itemRepository.FindByName）",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							nil,
							errors.New("itemRepository.FindByName"),
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					return m
				},
			},
			args: args{
				items: &request.Items{
					{
						Name:  "item1",
						Count: 1,
					},
				},
				accountID: int64(1),
				shardKey:  "SHARD_1",
				tx:        nil,
			},
			want:    nil,
			wantErr: errors.New("itemRepository.FindByName"),
		},
		{
			name: "正常：受け取りできる",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							&masterItemEntity.Item{
								ID:        1,
								Name:      "item1",
								Detail:    "detail1",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item1",
							"SHARD_1",
						).
						Return(
							nil,
							errors.New("itemBoxRepository.FindOrNilByAccountIDAndItemName"),
						)
					return m
				},
			},
			args: args{
				items: &request.Items{
					{
						Name:  "item1",
						Count: 1,
					},
				},
				accountID: int64(1),
				shardKey:  "SHARD_1",
				tx:        nil,
			},
			want:    nil,
			wantErr: errors.New("itemBoxRepository.FindOrNilByAccountIDAndItemName"),
		},
		{
			name: "正常：受け取りできる",
			fields: fields{
				itemRepository: func(ctrl *gomock.Controller) masterItemRepository.ItemRepository {
					m := masterItemRepository.NewMockItemRepository(ctrl)
					m.EXPECT().
						FindByName(
							"item1",
						).
						Return(
							&masterItemEntity.Item{
								ID:        1,
								Name:      "item1",
								Detail:    "detail1",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemBoxRepository: func(ctrl *gomock.Controller) userItemRepository.ItemBoxRepository {
					m := userItemRepository.NewMockItemBoxRepository(ctrl)
					m.EXPECT().
						FindOrNilByAccountIDAndItemName(
							int64(1),
							"item1",
							"SHARD_1",
						).
						Return(
							&userItemEntity.ItemBox{
								ID:        1,
								ShardKey:  "SHARD_1",
								AccountID: 1,
								ItemName:  "item1",
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							nil,
							errors.New("itemBoxRepository.Save"),
						)
					return m
				},
			},
			args: args{
				items: &request.Items{
					{
						Name:  "item1",
						Count: 1,
					},
				},
				accountID: int64(1),
				shardKey:  "SHARD_1",
				tx:        nil,
			},
			want:    nil,
			wantErr: errors.New("itemBoxRepository.Save"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				itemRepository:    tt.fields.itemRepository(ctrl),
				itemBoxRepository: tt.fields.itemBoxRepository(ctrl),
			}

			got, err := s.receiveItemInBox(tt.args.items, tt.args.accountID, tt.args.shardKey, tt.args.tx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ReceiveItemInBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceiveItemInBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
