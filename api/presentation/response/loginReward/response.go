package loginReward

import (
	masterLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/master/loginReward"
)

// ToItems アイテム一覧レスポンスを取得する
func ToItems(itemString string) (items Items, err error) {
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

// ToRewards 報酬一覧レスポンスを取得する
func ToRewards(lrrs *masterLoginRewardEntity.LoginRewardRewards) (rewards LoginRewardRewards, err error) {
	for _, lrr := range *lrrs {
		items, err := ToItems(lrrs.GetItems(lrr.StepNumber))
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
