package loginReward

import (
	"reflect"
	"testing"
	"time"
)

func TestLoginRewardRewardEntity_GetMaxStepNumber(t *testing.T) {
	type fields struct {
		rewards *LoginRewardRewards
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "正常：取得できる（リピートあり）",
			fields: fields{
				rewards: &LoginRewardRewards{
					{
						ID:                   1,
						LoginRewardModelName: "loginReward",
						Name:                 "reward1",
						StepNumber:           0,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:                   2,
						LoginRewardModelName: "loginReward",
						Name:                 "reward2",
						StepNumber:           1,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.rewards.GetMaxStepNumber()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaxStepNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginRewardRewardEntity_GetItems(t *testing.T) {
	type fields struct {
		rewards *LoginRewardRewards
	}
	type args struct {
		dayCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "正常：取得できる（リピートあり）",
			fields: fields{
				rewards: &LoginRewardRewards{
					{
						ID:                   1,
						LoginRewardModelName: "loginReward",
						Name:                 "reward1",
						StepNumber:           0,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:                   2,
						LoginRewardModelName: "loginReward",
						Name:                 "reward2",
						StepNumber:           1,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				dayCount: 2,
			},
			want: "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
		},
		{
			name: "正常：取得できる（リピートなし）",
			fields: fields{
				rewards: &LoginRewardRewards{
					{
						ID:                   1,
						LoginRewardModelName: "loginReward",
						Name:                 "reward1",
						StepNumber:           0,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:                   2,
						LoginRewardModelName: "loginReward",
						Name:                 "reward2",
						StepNumber:           1,
						Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":3}]",
						CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				dayCount: 1,
			},
			want: "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":3}]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.rewards.GetItems(tt.args.dayCount)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
