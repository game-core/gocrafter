package loginReward

// GetMaxStepNumber ステップナンバーの最大値を取得
func (es *LoginRewardRewards) GetMaxStepNumber() (maxStepNumber int) {
	for _, rewards := range *es {
		if rewards.StepNumber > maxStepNumber {
			maxStepNumber = rewards.StepNumber
		}
	}

	return maxStepNumber
}

// GetItemName アイテム名を取得
func (es *LoginRewardRewards) GetItemName(dayCount int) (itemName string) {
	maxStepNumber := es.GetMaxStepNumber()
	if maxStepNumber < dayCount && maxStepNumber > 0 {
		dayCount %= maxStepNumber
	}

	for _, rewards := range *es {
		if dayCount == rewards.StepNumber {
			itemName = rewards.ItemName
		}
	}

	return itemName
}

// GetItemCount アイテム数を取得
func (es *LoginRewardRewards) GetItemCount(dayCount int) (itemCount int) {
	maxStepNumber := es.GetMaxStepNumber()
	if maxStepNumber < dayCount && maxStepNumber > 0 {
		dayCount %= maxStepNumber
	}

	for _, rewards := range *es {
		if dayCount == rewards.StepNumber {
			itemCount = rewards.Count
		}
	}

	return itemCount
}
