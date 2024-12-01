package config

type Config struct {
	Sources      map[string]Integration `mapstructure:"sources"`
	Destinations map[string]Integration `mapstructure:"destinations"`
	Syncs        []Sync                 `mapstructure:"syncs"`
}

type Integration struct {
	Type   string                 `mapstructure:"type"`
	Extras map[string]interface{} `mapstructure:",remain"`
}

type Sync struct {
	Source      string `mapstructure:"source"`
	Destination string `mapstructure:"destination"`
}
