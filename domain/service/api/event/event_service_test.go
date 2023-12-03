package event

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/config/pointer"
	masterEventEntity "github.com/game-core/gocrafter/domain/entity/master/event"
	masterEventRepository "github.com/game-core/gocrafter/domain/repository/master/event"
)

func TestEventService_GetEventToEntity(t *testing.T) {
	type fields struct {
		eventRepository func(ctrl *gomock.Controller) masterEventRepository.EventRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *masterEventEntity.Event
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				eventRepository: func(ctrl *gomock.Controller) masterEventRepository.EventRepository {
					m := masterEventRepository.NewMockEventRepository(ctrl)
					m.EXPECT().
						FindByName(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: false,
								RepeatStartAt: nil,
								StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				name: "event",
			},
			want: &masterEventEntity.Event{
				ID:            1,
				Name:          "event",
				ResetHour:     9,
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
				EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
				CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（eventRepository.FindByName）",
			fields: fields{
				eventRepository: func(ctrl *gomock.Controller) masterEventRepository.EventRepository {
					m := masterEventRepository.NewMockEventRepository(ctrl)
					m.EXPECT().
						FindByName(
							"event",
						).
						Return(
							nil,
							errors.New("eventRepository.FindByName"),
						)
					return m
				},
			},
			args: args{
				name: "event",
			},
			want:    nil,
			wantErr: errors.New("eventRepository.FindByName"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &eventService{
				eventRepository: tt.fields.eventRepository(ctrl),
			}

			got, err := s.GetEventToEntity(tt.args.name)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetEventToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
