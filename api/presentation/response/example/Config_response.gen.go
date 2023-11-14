package example

type Config struct {
	GetConfig int64 `json:"get_config"`

	SetConfig int64 `json:"set_config"`
}

func ConfigResponse(getConfig int64, setConfig int64) *Config {
	return &Config{GetConfig: getConfig, SetConfig: setConfig}
}
