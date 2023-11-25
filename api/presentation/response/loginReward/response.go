package loginReward

import (
	masterLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/master/loginReward"
)

// GetItemResponses アイテム一覧レスポンスを取得する
func GetItemResponses(itemString string) (items Items, err error) {
	rewardItems := &masterLoginRewardEntity.LoginRewardItems{}
	if err := rewardItems.ToEntities(itemString); err != nil {
		return nil, err
	}

	for _, ri := range *rewardItems {
		item := Item{
			Name:  ri.Name,
			Count: ri.Count,
		}
		items = append(items, item)
	}

	return items, nil
}

// GetRewardResponses 報酬一覧レスポンスを取得する
func GetRewardResponses(lrrs *masterLoginRewardEntity.LoginRewardRewards) (rewards LoginRewardRewards, err error) {
	for _, lrr := range *lrrs {
		items, err := GetItemResponses(lrrs.GetItems(lrr.StepNumber))
		if err != nil {
			return nil, err
		}

		reward := &LoginRewardReward{
			ID:         lrr.ID,
			Name:       lrr.Name,
			Items:      items,
			StepNumber: lrr.StepNumber,
		}
		rewards = append(rewards, *reward)
	}

	return rewards, nil
}
