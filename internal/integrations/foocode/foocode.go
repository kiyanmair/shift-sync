package foocode

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
	var i FooCode
	if err := mapstructure.Decode(cfg.Extras, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *FooCode) Validate(direction core.Direction) error {
	if i.APIKey == "" {
		return errors.New("api_key cannot be empty")
	}
	if i.TeamName == "" {
		return errors.New("team_name cannot be empty")
	}
	if direction == core.DestinationDirection {
		if i.Role == "" {
			return errors.New("role cannot be empty")
		}
	}
	return nil
}

func (i *FooCode) GetUsers() ([]string, error) {
	return []string{"foocode.user@example.com"}, nil
}

func (i *FooCode) SetUsers(users []string) error {
	return nil
}
