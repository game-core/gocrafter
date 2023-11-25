package loginReward

import (
	"encoding/json"
	"errors"
)

func (e *LoginRewardItems) ToJson() (string, error) {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string(jsonData), nil
}

func (e *LoginRewardItems) ToEntities(string string) error {
	err := json.Unmarshal([]byte(string), e)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
