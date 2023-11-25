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

// GetItems アイテムを取得
func (es *LoginRewardRewards) GetItems(dayCount int) (items string) {
	maxStepNumber := es.GetMaxStepNumber()
	if maxStepNumber < dayCount && maxStepNumber > 0 {
		dayCount %= maxStepNumber
	}

	for _, rewards := range *es {
		if dayCount == rewards.StepNumber {
			items = rewards.Items
		}
	}

	return items
}
