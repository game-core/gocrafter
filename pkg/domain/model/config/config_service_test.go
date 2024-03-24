package config

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/config/masterConfig"
)

func TestNewConfigService_NewConfigService(t *testing.T) {
	type args struct {
		masterConfigMysqlRepository masterConfig.MasterConfigMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ConfigService
	}{
		{
			name: "正常",
			args: args{
				masterConfigMysqlRepository: nil,
			},
			want: &configService{
				masterConfigMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConfigService(
				tt.args.masterConfigMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigService_GetAll(t *testing.T) {
	type fields struct {
		masterConfigMysqlRepository func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    masterConfig.MasterConfigs
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterConfigMysqlRepository: func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository {
					m := masterConfig.NewMockMasterConfigMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterConfig.MasterConfigs{
								{
									Id:         1,
									Name:       "ルーム最大数",
									ConfigType: enum.ConfigType_MaxRoom,
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
			want: masterConfig.MasterConfigs{
				{
					Id:         1,
					Name:       "ルーム最大数",
					ConfigType: enum.ConfigType_MaxRoom,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterConfigMysqlRepository.FindList",
			fields: fields{
				masterConfigMysqlRepository: func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository {
					m := masterConfig.NewMockMasterConfigMysqlRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterConfigMysqlRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &configService{
				masterConfigMysqlRepository: tt.fields.masterConfigMysqlRepository(ctrl),
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

func TestConfigService_GetByConfigType(t *testing.T) {
	type fields struct {
		masterConfigMysqlRepository func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository
	}
	type args struct {
		ctx        context.Context
		configType enum.ConfigType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *masterConfig.MasterConfig
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterConfigMysqlRepository: func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository {
					m := masterConfig.NewMockMasterConfigMysqlRepository(ctrl)
					m.EXPECT().
						FindByConfigType(
							nil,
							enum.ConfigType_MaxRoom,
						).
						Return(
							&masterConfig.MasterConfig{
								Id:         1,
								Name:       "ルーム最大数",
								ConfigType: enum.ConfigType_MaxRoom,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: &masterConfig.MasterConfig{
				Id:         1,
				Name:       "ルーム最大数",
				ConfigType: enum.ConfigType_MaxRoom,
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterConfigMysqlRepository.FindByConfigType",
			fields: fields{
				masterConfigMysqlRepository: func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository {
					m := masterConfig.NewMockMasterConfigMysqlRepository(ctrl)
					m.EXPECT().
						FindByConfigType(
							nil,
							enum.ConfigType_MaxRoom,
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
			wantErr: errors.NewMethodError("s.masterConfigMysqlRepository.FindByConfigType", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &configService{
				masterConfigMysqlRepository: tt.fields.masterConfigMysqlRepository(ctrl),
			}

			got, err := s.GetByConfigType(tt.args.ctx, tt.args.configType)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetByConfigType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByConfigType() = %v, want %v", got, tt.want)
			}
		})
	}
}
