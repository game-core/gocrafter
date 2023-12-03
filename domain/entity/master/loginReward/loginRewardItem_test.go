package loginReward

import (
	"reflect"
	"testing"
)

func TestLoginRewardItemEntity_ToJson(t *testing.T) {
	type fields struct {
		items *LoginRewardItems
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				items: &LoginRewardItems{
					{
						Name:  "name1",
						Count: 1,
					},
					{
						Name:  "name2",
						Count: 2,
					},
				},
			},
			want: "[{\"id\":0,\"name\":\"name1\",\"count\":1,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"},{\"id\":0,\"name\":\"name2\",\"count\":2,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.items.ToJson()
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

func TestLoginRewardItemEntity_ToEntities(t *testing.T) {
	type args struct {
		items string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "正常：取得できる",
			args: args{
				items: "[{\"id\":0,\"name\":\"name1\",\"count\":1,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"},{\"id\":0,\"name\":\"name2\",\"count\":2,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			items := &LoginRewardItems{}

			err := items.ToEntities(tt.args.items)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetEventToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
