package loginReward

import (
	"time"
)

type LoginRewardItems []LoginRewardItem

type LoginRewardItem struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *LoginRewardItem) TableName() string {
	return "login_reward_item"
}
