package foochat

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterIntegration("foochat", NewFooChat)
}

type FooChat struct {
	Token   string   `mapstructure:"token"`
	GroupID string   `mapstructure:"group_id"`
	Include []string `mapstructure:"include"`
}

func NewFooChat(cfg config.Integration) (core.Integration, error) {
	var i FooChat
	if err := mapstructure.Decode(cfg.Extras, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *FooChat) Validate(direction core.IntegrationDirection) error {
	if i.Token == "" {
		return errors.New("token cannot be empty")
	}
	if i.GroupID == "" {
		return errors.New("group_id cannot be empty")
	}
	if direction == core.SourceDirection {
		if len(i.Include) == 0 {
			return errors.New("include cannot be empty")
		}
	}
	return nil
}

func (i *FooChat) GetUsers() ([]string, error) {
	return []string{"foochat.user@example.com"}, nil
}

func (i *FooChat) SetUsers(users []string) error {
	return nil
}
