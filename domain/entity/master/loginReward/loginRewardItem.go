package loginReward

import (
	"encoding/json"
)

func (e *LoginRewardItems) ToJson() (string, error) {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (e *LoginRewardItems) ToEntities(string string) error {
	err := json.Unmarshal([]byte(string), e)
	if err != nil {
		return err
	}

	return nil
}
