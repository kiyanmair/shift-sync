package config

type Config struct {
	Sources      map[string]Source      `mapstructure:"sources"`
	Destinations map[string]Destination `mapstructure:"destinations"`
	Syncs        []Sync                 `mapstructure:"syncs"`
}

type Source struct {
	Type   string                 `mapstructure:"type"`
	Extras map[string]interface{} `mapstructure:",remain"`
}

type Destination struct {
	Type   string                 `mapstructure:"type"`
	Extras map[string]interface{} `mapstructure:",remain"`
}

type Sync struct {
	Source      string `mapstructure:"source"`
	Destination string `mapstructure:"destination"`
}
