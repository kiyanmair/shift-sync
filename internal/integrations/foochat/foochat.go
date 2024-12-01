package destination

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
	var dest FooChat
	if err := mapstructure.Decode(cfg.Extras, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

func (i *FooChat) Valid() (bool, error) {
	if i.Token == "" {
		return false, errors.New("token cannot be empty")
	}
	if i.GroupID == "" {
		return false, errors.New("group_id cannot be empty")
	}
	if core.IsSource(i) && len(i.Include) == 0 {
		return false, errors.New("include cannot be empty")
	}
	return true, nil
}

func (i *FooChat) GetUsers() ([]string, error) {
	return []string{"foochat.user@example.com"}, nil
}

func (i *FooChat) SetUsers(users []string) error {
	return nil
}
