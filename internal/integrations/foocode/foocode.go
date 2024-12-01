package source

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterIntegration("foocode", NewFooCode)
}

type FooCode struct {
	APIKey   string `mapstructure:"api_key"`
	TeamName string `mapstructure:"team_name"`
	Role     string `mapstructure:"role"`
}

func NewFooCode(cfg config.Integration) (core.Integration, error) {
	var src FooCode
	if err := mapstructure.Decode(cfg.Extras, &src); err != nil {
		return nil, err
	}
	return &src, nil
}

func (i *FooCode) Valid(direction core.IntegrationDirection) (bool, error) {
	if i.APIKey == "" {
		return false, errors.New("api_key cannot be empty")
	}
	if i.TeamName == "" {
		return false, errors.New("schedule_id cannot be empty")
	}
	if direction == core.DestinationDirection {
		if i.Role == "" {
			return false, errors.New("role cannot be empty")
		}
	}
	return true, nil
}

func (i *FooCode) GetUsers() ([]string, error) {
	return []string{"foocode.user@example.com"}, nil
}

func (i *FooCode) SetUsers(users []string) error {
	return nil
}
