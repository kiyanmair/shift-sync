package config

type Config struct {
	Sources      []Source      `yaml:"sources"`
	Destinations []Destination `yaml:"destinations"`
	Syncs        []Sync        `yaml:"syncs"`
}

type Source struct {
	ID      string                 `yaml:"id"`
	Type    string                 `yaml:"type"`
	Details map[string]interface{} `yaml:"details"`
}

type Destination struct {
	ID      string                 `yaml:"id"`
	Type    string                 `yaml:"type"`
	Details map[string]interface{} `yaml:"details"`
}

type Sync struct {
	ID            string `yaml:"id"`
	DestinationID string `yaml:"destination_id"`
	SourceID      string `yaml:"source_id"`
}
