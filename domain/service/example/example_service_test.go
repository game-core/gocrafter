package example

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	response "github.com/game-core/gocrafter/api/presentation/response/example"
	"github.com/game-core/gocrafter/config/pointer"
	exampleEntity "github.com/game-core/gocrafter/domain/entity/master/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
)

func TestExampleService_ListExample(t *testing.T) {
	type fields struct {
		exampleRepository func(ctrl *gomock.Controller) exampleRepository.ExampleRepository
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ListExample
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				exampleRepository: func(ctrl *gomock.Controller) exampleRepository.ExampleRepository {
					m := exampleRepository.NewMockExampleRepository(ctrl)
					m.EXPECT().
						List(
							10,
						).
						Return(
							&exampleEntity.Examples{
								{
									Name:      "name1",
									Detail:    pointer.StringToPointer("detail1"),
									Count:     1,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									Name:      "name2",
									Detail:    pointer.StringToPointer("detail2"),
									Count:     2,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									Name:      "name3",
									Detail:    pointer.StringToPointer("detail3"),
									Count:     3,
									CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				limit: 10,
			},
			want: &response.ListExample{
				Status: 200,
				Items: &response.Examples{
					{
						Name:   "name1",
						Detail: pointer.StringToPointer("detail1"),
						Count:  1,
					},
					{
						Name:   "name2",
						Detail: pointer.StringToPointer("detail2"),
						Count:  2,
					},
					{
						Name:   "name3",
						Detail: pointer.StringToPointer("detail3"),
						Count:  3,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（exampleRepository.List）",
			fields: fields{
				exampleRepository: func(ctrl *gomock.Controller) exampleRepository.ExampleRepository {
					m := exampleRepository.NewMockExampleRepository(ctrl)
					m.EXPECT().
						List(
							10,
						).
						Return(
							nil,
							errors.New("exampleRepository.List"),
						)
					return m
				},
			},
			args: args{
				limit: 10,
			},
			want:    nil,
			wantErr: errors.New("exampleRepository.List"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &exampleService{
				exampleRepository: tt.fields.exampleRepository(ctrl),
			}

			got, err := s.ListExample(tt.args.limit)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ListExample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
