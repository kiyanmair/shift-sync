package source

import (
	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterSource("example", NewExampleSource)
}

type ExampleSource struct {
	APIKey     string `mapstructure:"api_key"`
	ScheduleID string `mapstructure:"schedule_id"`
}

func NewExampleSource(cfg config.Source) (core.Source, error) {
	var src ExampleSource
	if err := mapstructure.Decode(cfg.Extras, &src); err != nil {
		return nil, err
	}
	return &src, nil
}

func (s *ExampleSource) FetchUsers() ([]string, error) {
	return []string{"user@example.com"}, nil
}
