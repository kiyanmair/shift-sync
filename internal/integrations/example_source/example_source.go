package source

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterIntegration("example_source", NewExampleSource)
}

type ExampleSource struct {
	APIKey     string `mapstructure:"api_key"`
	ScheduleID string `mapstructure:"schedule_id"`
}

func NewExampleSource(cfg config.Integration) (core.Integration, error) {
	var src ExampleSource
	if err := mapstructure.Decode(cfg.Extras, &src); err != nil {
		return nil, err
	}
	return &src, nil
}

func (s *ExampleSource) Valid() (bool, error) {
	if s.APIKey == "" {
		return false, errors.New("api_key cannot be empty")
	}
	if s.ScheduleID == "" {
		return false, errors.New("schedule_id cannot be empty")
	}
	return true, nil
}

func (s *ExampleSource) GetUsers() ([]string, error) {
	return []string{"user@example.com"}, nil
}
