// Package commonRankingWorld ワールドランキング
package commonRankingWorld

import (
	"encoding/json"

	"time"
)

type CommonRankingWorlds []*CommonRankingWorld

type CommonRankingWorld struct {
	MasterRankingId int64
	UserId          string
	Score           int32
	RankedAt        time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCommonRankingWorld() *CommonRankingWorld {
	return &CommonRankingWorld{}
}

func NewCommonRankingWorlds() CommonRankingWorlds {
	return CommonRankingWorlds{}
}

func SetCommonRankingWorld(masterRankingId int64, userId string, score int32, rankedAt time.Time, createdAt time.Time, updatedAt time.Time) *CommonRankingWorld {
	return &CommonRankingWorld{
		MasterRankingId: masterRankingId,
		UserId:          userId,
		Score:           score,
		RankedAt:        rankedAt,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (t *CommonRankingWorld) TableToJson() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (t *CommonRankingWorld) JsonToTable(data string) error {
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return err
	}

	return nil
}

func (t *CommonRankingWorld) TableName() string {
	return "common_ranking_world"
}
