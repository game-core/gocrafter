// Package commonRankingRoom ルームランキング
package commonRankingRoom

import (
	"encoding/json"

	"time"
)

type CommonRankingRooms []*CommonRankingRoom

type CommonRankingRoom struct {
	MasterRankingId int64
	RoomId          string
	UserId          string
	Score           int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCommonRankingRoom() *CommonRankingRoom {
	return &CommonRankingRoom{}
}

func NewCommonRankingRooms() CommonRankingRooms {
	return CommonRankingRooms{}
}

func SetCommonRankingRoom(masterRankingId int64, roomId string, userId string, score int32, createdAt time.Time, updatedAt time.Time) *CommonRankingRoom {
	return &CommonRankingRoom{
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
		UserId:          userId,
		Score:           score,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (t *CommonRankingRoom) TableToJson() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (t *CommonRankingRoom) JsonToTable(data string) error {
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return err
	}

	return nil
}

func (t *CommonRankingRoom) TableName() string {
	return "common_ranking_room"
}
