package shard

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	response "github.com/game-core/gocrafter/api/presentation/response/shard"
	shardEntity "github.com/game-core/gocrafter/domain/entity/config/shard"
	configRepository "github.com/game-core/gocrafter/domain/repository/config"
	shardRepository "github.com/game-core/gocrafter/domain/repository/config/shard"
)

func TestShardService_GetShard(t *testing.T) {
	type fields struct {
		transactionRepository func(ctrl *gomock.Controller) configRepository.TransactionRepository
		shardRepository       func(ctrl *gomock.Controller) shardRepository.ShardRepository
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetShard
		wantErr error
	}{
		{
			name: "正常：取得できる（切り替えなし）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							nil,
						).
						Return(
							nil,
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					m.EXPECT().
						List(
							64,
						).
						Return(
							&shardEntity.Shards{
								{
									ID:        1,
									ShardKey:  "SHARD_1",
									Name:      "name1",
									Count:     1,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:        2,
									ShardKey:  "SHARD_2",
									Name:      "name2",
									Count:     2,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Save(
							&shardEntity.Shard{
								ID:        1,
								ShardKey:  "SHARD_1",
								Name:      "name1",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						).
						Return(
							&shardEntity.Shard{
								ID:        1,
								ShardKey:  "SHARD_2",
								Name:      "name1",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			want: &response.GetShard{
				Status:       200,
				NextShardKey: "SHARD_1",
				Shards: &response.Shards{
					{
						ID:       1,
						ShardKey: "SHARD_1",
						Name:     "name1",
						Count:    1,
					},
					{
						ID:       2,
						ShardKey: "SHARD_2",
						Name:     "name2",
						Count:    2,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：取得できる（切り替えあり）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							nil,
						).
						Return(
							nil,
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					m.EXPECT().
						List(
							64,
						).
						Return(
							&shardEntity.Shards{
								{
									ID:        1,
									ShardKey:  "SHARD_1",
									Name:      "name1",
									Count:     2,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:        2,
									ShardKey:  "SHARD_2",
									Name:      "name2",
									Count:     1,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Save(
							&shardEntity.Shard{
								ID:        2,
								ShardKey:  "SHARD_2",
								Name:      "name2",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						).
						Return(
							&shardEntity.Shard{
								ID:        2,
								ShardKey:  "SHARD_2",
								Name:      "name2",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			want: &response.GetShard{
				Status:       200,
				NextShardKey: "SHARD_2",
				Shards: &response.Shards{
					{
						ID:       1,
						ShardKey: "SHARD_1",
						Name:     "name1",
						Count:    2,
					},
					{
						ID:       2,
						ShardKey: "SHARD_2",
						Name:     "name2",
						Count:    1,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（transactionRepository.Begin）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							errors.New("transactionRepository.Begin"),
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					return m
				},
			},
			want:    nil,
			wantErr: errors.New("transactionRepository.Begin"),
		},
		{
			name: "異常：エラー（shardRepository.List）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					m.EXPECT().
						List(
							64,
						).
						Return(
							nil,
							errors.New("shardRepository.List"),
						)
					return m
				},
			},
			want:    nil,
			wantErr: errors.New("shardRepository.List"),
		},
		{
			name: "異常：エラー（failed to shardRepository.List）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					m.EXPECT().
						List(
							64,
						).
						Return(
							&shardEntity.Shards{},
							nil,
						)
					return m
				},
			},
			want:    nil,
			wantErr: errors.New("failed to shardRepository.List"),
		},
		{
			name: "異常：エラー（hardRepository.Update）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) configRepository.TransactionRepository {
					m := configRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				shardRepository: func(ctrl *gomock.Controller) shardRepository.ShardRepository {
					m := shardRepository.NewMockShardRepository(ctrl)
					m.EXPECT().
						List(
							64,
						).
						Return(
							&shardEntity.Shards{
								{
									ID:        1,
									ShardKey:  "SHARD_1",
									Name:      "name1",
									Count:     1,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:        2,
									ShardKey:  "SHARD_2",
									Name:      "name2",
									Count:     2,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Save(
							&shardEntity.Shard{
								ID:        1,
								ShardKey:  "SHARD_1",
								Name:      "name1",
								Count:     2,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						).
						Return(
							nil,
							errors.New("hardRepository.Update"),
						)
					return m
				},
			},
			want:    nil,
			wantErr: errors.New("hardRepository.Update"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &shardService{
				transactionRepository: tt.fields.transactionRepository(ctrl),
				shardRepository:       tt.fields.shardRepository(ctrl),
			}

			got, err := s.GetShard()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ListShard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListShard() = %v, want %v", got, tt.want)
			}
		})
	}
}
