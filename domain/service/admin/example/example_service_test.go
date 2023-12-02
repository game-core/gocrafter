package example

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/admin/presentation/request/example"
	response "github.com/game-core/gocrafter/admin/presentation/response/example"
	"github.com/game-core/gocrafter/config/pointer"
	exampleEntity "github.com/game-core/gocrafter/domain/entity/admin/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/admin/example"
)

func TestExampleService_GetExample(t *testing.T) {
	type fields struct {
		exampleRepository func(ctrl *gomock.Controller) exampleRepository.ExampleRepository
	}
	type args struct {
		req *request.GetExample
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetExample
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				exampleRepository: func(ctrl *gomock.Controller) exampleRepository.ExampleRepository {
					m := exampleRepository.NewMockExampleRepository(ctrl)
					m.EXPECT().
						FindByID(
							int64(1),
						).
						Return(
							&exampleEntity.Example{
								Name:      "name1",
								Detail:    pointer.StringToPointer("detail1"),
								Count:     1,
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.GetExample{
					ID: 1,
				},
			},
			want: &response.GetExample{
				Status: 200,
				Example: &response.Example{
					Name:   "name1",
					Detail: pointer.StringToPointer("detail1"),
					Count:  1,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（exampleRepository.FindByID）",
			fields: fields{
				exampleRepository: func(ctrl *gomock.Controller) exampleRepository.ExampleRepository {
					m := exampleRepository.NewMockExampleRepository(ctrl)
					m.EXPECT().
						FindByID(
							int64(1),
						).
						Return(
							nil,
							errors.New("exampleRepository.FindByID"),
						)
					return m
				},
			},
			args: args{
				req: &request.GetExample{
					ID: 1,
				},
			},
			want:    nil,
			wantErr: errors.New("exampleRepository.FindByID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &exampleService{
				exampleRepository: tt.fields.exampleRepository(ctrl),
			}

			got, err := s.GetExample(tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetExample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
