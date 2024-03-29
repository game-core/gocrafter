package event

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/event/masterEvent"
)

func TestNewEventService_NewEventService(t *testing.T) {
	type args struct {
		masterEventMysqlRepository masterEvent.MasterEventMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want EventService
	}{
		{
			name: "正常",
			args: args{
				masterEventMysqlRepository: nil,
			},
			want: &eventService{
				masterEventMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEventService(
				tt.args.masterEventMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventService_Get(t *testing.T) {
	type fields struct {
		masterEventMysqlRepository func(ctrl *gomock.Controller) masterEvent.MasterEventMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *EventGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *EventGetResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterEventMysqlRepository: func(ctrl *gomock.Controller) masterEvent.MasterEventMysqlRepository {
					m := masterEvent.NewMockMasterEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterEvent.MasterEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     1,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &EventGetRequest{
					EventId: 1,
				},
			},
			want: &EventGetResponse{
				MasterEvent: &masterEvent.MasterEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     1,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterEventMysqlRepository.Find",
			fields: fields{
				masterEventMysqlRepository: func(ctrl *gomock.Controller) masterEvent.MasterEventMysqlRepository {
					m := masterEvent.NewMockMasterEventMysqlRepository(ctrl)
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
				req: &EventGetRequest{
					EventId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterEventMysqlRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &eventService{
				masterEventMysqlRepository: tt.fields.masterEventMysqlRepository(ctrl),
			}

			got, err := s.Get(tt.args.ctx, tt.args.req)
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
