package loginReward

import (
	"time"
)

type LoginRewardModels []LoginRewardModel

type LoginRewardModel struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	EventName string `json:"event_name"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *LoginRewardModel) TableName() string {
	return "login_reward_model"
}
