package shard

import (
	"context"
	"reflect"
	"testing"

	"gorm.io/gorm"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/shard/commonShard"
)

func TestShardService_NewShardService(t *testing.T) {
	type args struct {
		commonShardRepository commonShard.CommonShardRepository
	}
	tests := []struct {
		name string
		args args
		want ShardService
	}{
		{
			name: "正常",
			args: args{
				commonShardRepository: nil,
			},
			want: &shardService{
				commonShardRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewShardService(
				tt.args.commonShardRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShardService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShardService_GetShardKey(t *testing.T) {
	type fields struct {
		commonShardRepository func(ctrl *gomock.Controller) commonShard.CommonShardRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				commonShardRepository: func(ctrl *gomock.Controller) commonShard.CommonShardRepository {
					m := commonShard.NewMockCommonShardRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							commonShard.CommonShards{
								{
									Id:       1,
									ShardKey: "SHARD_1",
									Name:     "name1",
									Count:    1,
								},
								{
									Id:       2,
									ShardKey: "SHARD_2",
									Name:     "name2",
									Count:    1,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
			},
			want:    "SHARD_1",
			wantErr: nil,
		},
		{
			name: "異常：s.commonShardRepository.FindList",
			fields: fields{
				commonShardRepository: func(ctrl *gomock.Controller) commonShard.CommonShardRepository {
					m := commonShard.NewMockCommonShardRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
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
			},
			want:    "",
			wantErr: errors.NewMethodError("shards.GetShardKey: failed to s.commonShardRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：common_shard does not exist",
			fields: fields{
				commonShardRepository: func(ctrl *gomock.Controller) commonShard.CommonShardRepository {
					m := commonShard.NewMockCommonShardRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							commonShard.CommonShards{},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
			},
			want:    "",
			wantErr: errors.NewError("failed to shards.GetShardKey: common_shard does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commonShard.CommonShardInstances = commonShard.NewCommonShards()
			ctrl := gomock.NewController(t)

			s := &shardService{
				commonShardRepository: tt.fields.commonShardRepository(ctrl),
			}

			got, err := s.GetShardKey(tt.args.ctx, tt.args.tx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetShardKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShardKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
