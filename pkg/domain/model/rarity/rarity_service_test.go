package rarity

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity/masterRarity"
)

func TestNewRarityService_NewRarityService(t *testing.T) {
	type args struct {
		masterRarityMysqlRepository masterRarity.MasterRarityMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want RarityService
	}{
		{
			name: "正常",
			args: args{
				masterRarityMysqlRepository: nil,
			},
			want: &rarityService{
				masterRarityMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRarityService(
				tt.args.masterRarityMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRarityService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRarityService_GetAll(t *testing.T) {
	type fields struct {
		masterRarityMysqlRepository func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    masterRarity.MasterRarities
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterRarityMysqlRepository: func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository {
					m := masterRarity.NewMockMasterRarityMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterRarity.MasterRarities{
								{
									Id:         1,
									Name:       "ノーマル",
									RarityType: enum.RarityType_N,
								},
								{
									Id:         2,
									Name:       "レア",
									RarityType: enum.RarityType_R,
								},
								{
									Id:         3,
									Name:       "スーパーレア",
									RarityType: enum.RarityType_SR,
								},
								{
									Id:         4,
									Name:       "スーパースペシャルレア",
									RarityType: enum.RarityType_SSR,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: masterRarity.MasterRarities{
				{
					Id:         1,
					Name:       "ノーマル",
					RarityType: enum.RarityType_N,
				},
				{
					Id:         2,
					Name:       "レア",
					RarityType: enum.RarityType_R,
				},
				{
					Id:         3,
					Name:       "スーパーレア",
					RarityType: enum.RarityType_SR,
				},
				{
					Id:         4,
					Name:       "スーパースペシャルレア",
					RarityType: enum.RarityType_SSR,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterRarityMysqlRepository.FindList",
			fields: fields{
				masterRarityMysqlRepository: func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository {
					m := masterRarity.NewMockMasterRarityMysqlRepository(ctrl)
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
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRarityMysqlRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &rarityService{
				masterRarityMysqlRepository: tt.fields.masterRarityMysqlRepository(ctrl),
			}

			got, err := s.GetAll(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRarityService_GetByRarityType(t *testing.T) {
	type fields struct {
		masterRarityMysqlRepository func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository
	}
	type args struct {
		ctx        context.Context
		rarityType enum.RarityType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *masterRarity.MasterRarity
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterRarityMysqlRepository: func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository {
					m := masterRarity.NewMockMasterRarityMysqlRepository(ctrl)
					m.EXPECT().
						FindByRarityType(
							nil,
							enum.RarityType_N,
						).
						Return(
							&masterRarity.MasterRarity{
								Id:         1,
								Name:       "ノーマル",
								RarityType: enum.RarityType_N,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: &masterRarity.MasterRarity{
				Id:         1,
				Name:       "ノーマル",
				RarityType: enum.RarityType_N,
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterRarityMysqlRepository.FindByRarityType",
			fields: fields{
				masterRarityMysqlRepository: func(ctrl *gomock.Controller) masterRarity.MasterRarityMysqlRepository {
					m := masterRarity.NewMockMasterRarityMysqlRepository(ctrl)
					m.EXPECT().
						FindByRarityType(
							nil,
							enum.RarityType_N,
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
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRarityMysqlRepository.FindByRarityType", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &rarityService{
				masterRarityMysqlRepository: tt.fields.masterRarityMysqlRepository(ctrl),
			}

			got, err := s.GetByRarityType(tt.args.ctx, tt.args.rarityType)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetByRarityType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByRarityType() = %v, want %v", got, tt.want)
			}
		})
	}
}
