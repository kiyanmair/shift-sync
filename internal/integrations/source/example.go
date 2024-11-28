package source

import (
	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterSource("example", NewExampleSource)
}

type ExampleSource struct {
	ID         string
	APIKey     string `mapstructure:"api_key"`
	ScheduleID string `mapstructure:"schedule_id"`
}

func NewExampleSource(cfg config.Source) (core.Source, error) {
	var source ExampleSource
	source.ID = cfg.ID

	if err := mapstructure.Decode(cfg.Details, &source); err != nil {
		return nil, err
	}

	return &source, nil
}

func (s *ExampleSource) FetchUsers() ([]string, error) {
	return []string{"user@example.com"}, nil
}
