package loginReward

// GetMaxStepNumber ステップナンバーの最大値を取得
func (es *LoginRewardRewards) GetMaxStepNumber() int {
	var maxStepNumber int
	for _, rewards := range *es {
		if rewards.StepNumber > maxStepNumber {
			maxStepNumber = rewards.StepNumber
		}
	}

	return maxStepNumber
}

// GetItemName アイテム名を取得
func (es *LoginRewardRewards) GetItemName(dayCount int) string {
	maxStepNumber := es.GetMaxStepNumber()
	if maxStepNumber < dayCount && maxStepNumber > 0 {
		dayCount %= maxStepNumber
	}

	var itemName string
	for _, rewards := range *es {
		if dayCount == rewards.StepNumber {
			itemName = rewards.ItemName
		}
	}

	return itemName
}
