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
									ConfigType: enum.ConfigType_MaxRoomNumber,
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
					ConfigType: enum.ConfigType_MaxRoomNumber,
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

func TestConfigService_Get(t *testing.T) {
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
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							&masterConfig.MasterConfig{
								Id:         1,
								Name:       "ルーム最大数",
								ConfigType: enum.ConfigType_MaxRoomNumber,
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
				ConfigType: enum.ConfigType_MaxRoomNumber,
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
							enum.ConfigType_MaxRoomNumber,
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

			got, err := s.Get(tt.args.ctx, tt.args.configType)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigService_GetInt32(t *testing.T) {
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
		want    int32
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
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							&masterConfig.MasterConfig{
								Id:         1,
								Name:       "ルーム最大数",
								ConfigType: enum.ConfigType_MaxRoomNumber,
								Value:      "1",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    1,
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
							enum.ConfigType_MaxRoomNumber,
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
			want:    0,
			wantErr: errors.NewMethodError("s.masterConfigMysqlRepository.FindByConfigType", errors.NewTestError()),
		},
		{
			name: "異常：changes.StringToInt32",
			fields: fields{
				masterConfigMysqlRepository: func(ctrl *gomock.Controller) masterConfig.MasterConfigMysqlRepository {
					m := masterConfig.NewMockMasterConfigMysqlRepository(ctrl)
					m.EXPECT().
						FindByConfigType(
							nil,
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							&masterConfig.MasterConfig{
								Id:         1,
								Name:       "ルーム最大数",
								ConfigType: enum.ConfigType_MaxRoomNumber,
								Value:      "a",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    0,
			wantErr: errors.NewError("failed to changes.StringToInt32: strconv.Atoi: parsing \"a\": invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &configService{
				masterConfigMysqlRepository: tt.fields.masterConfigMysqlRepository(ctrl),
			}

			got, err := s.GetInt32(tt.args.ctx, tt.args.configType)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
