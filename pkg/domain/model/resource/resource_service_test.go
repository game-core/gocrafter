package resource

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/resource/masterResource"
)

func TestNewResourceService_NewResourceService(t *testing.T) {
	type args struct {
		masterResourceRepository masterResource.MasterResourceRepository
	}
	tests := []struct {
		name string
		args args
		want ResourceService
	}{
		{
			name: "正常",
			args: args{
				masterResourceRepository: nil,
			},
			want: &resourceService{
				masterResourceRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewResourceService(
				tt.args.masterResourceRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceService_GetAll(t *testing.T) {
	type fields struct {
		masterResourceRepository func(ctrl *gomock.Controller) masterResource.MasterResourceRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    masterResource.MasterResources
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterResourceRepository: func(ctrl *gomock.Controller) masterResource.MasterResourceRepository {
					m := masterResource.NewMockMasterResourceRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterResource.MasterResources{
								{
									Id:           1,
									Name:         "ノーマル",
									ResourceType: enum.ResourceType_Normal,
								},
								{
									Id:           2,
									Name:         "カード",
									ResourceType: enum.ResourceType_Card,
								},
								{
									Id:           3,
									Name:         "チケット",
									ResourceType: enum.ResourceType_Ticket,
								},
								{
									Id:           4,
									Name:         "コイン",
									ResourceType: enum.ResourceType_Coin,
								},
								{
									Id:           5,
									Name:         "クリスタル",
									ResourceType: enum.ResourceType_Crystal,
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
			want: masterResource.MasterResources{
				{
					Id:           1,
					Name:         "ノーマル",
					ResourceType: enum.ResourceType_Normal,
				},
				{
					Id:           2,
					Name:         "カード",
					ResourceType: enum.ResourceType_Card,
				},
				{
					Id:           3,
					Name:         "チケット",
					ResourceType: enum.ResourceType_Ticket,
				},
				{
					Id:           4,
					Name:         "コイン",
					ResourceType: enum.ResourceType_Coin,
				},
				{
					Id:           5,
					Name:         "クリスタル",
					ResourceType: enum.ResourceType_Crystal,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterResourceRepository.FindList",
			fields: fields{
				masterResourceRepository: func(ctrl *gomock.Controller) masterResource.MasterResourceRepository {
					m := masterResource.NewMockMasterResourceRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterResourceRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &resourceService{
				masterResourceRepository: tt.fields.masterResourceRepository(ctrl),
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

func TestResourceService_GetByResourceType(t *testing.T) {
	type fields struct {
		masterResourceRepository func(ctrl *gomock.Controller) masterResource.MasterResourceRepository
	}
	type args struct {
		ctx          context.Context
		resourceType enum.ResourceType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *masterResource.MasterResource
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterResourceRepository: func(ctrl *gomock.Controller) masterResource.MasterResourceRepository {
					m := masterResource.NewMockMasterResourceRepository(ctrl)
					m.EXPECT().
						FindByResourceType(
							nil,
							enum.ResourceType_Normal,
						).
						Return(
							&masterResource.MasterResource{
								Id:           1,
								Name:         "ノーマル",
								ResourceType: enum.ResourceType_Normal,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: &masterResource.MasterResource{
				Id:           1,
				Name:         "ノーマル",
				ResourceType: enum.ResourceType_Normal,
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterResourceRepository.FindByResourceType",
			fields: fields{
				masterResourceRepository: func(ctrl *gomock.Controller) masterResource.MasterResourceRepository {
					m := masterResource.NewMockMasterResourceRepository(ctrl)
					m.EXPECT().
						FindByResourceType(
							nil,
							enum.ResourceType_Normal,
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
			wantErr: errors.NewMethodError("s.masterResourceRepository.FindByResourceType", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &resourceService{
				masterResourceRepository: tt.fields.masterResourceRepository(ctrl),
			}

			got, err := s.GetByResourceType(tt.args.ctx, tt.args.resourceType)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetByResourceType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByResourceType() = %v, want %v", got, tt.want)
			}
		})
	}
}
