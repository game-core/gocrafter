package example

type Config struct {
	GetConfig int64 `json:"get_config"`

	SetConfig int64 `json:"set_config"`
}

func ConfigResponse(setConfig int64, getConfig int64) *Config {
	return &Config{SetConfig: setConfig, GetConfig: getConfig}
}
