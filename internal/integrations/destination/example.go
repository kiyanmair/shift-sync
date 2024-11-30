package destination

import (
	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/internal/config"
	"github.com/kiyanmair/shift-sync/internal/core"
)

func init() {
	core.RegisterDestination("example", NewExampleDestination)
}

type ExampleDestination struct {
	Token   string `mapstructure:"token"`
	GroupID string `mapstructure:"group_id"`
}

func NewExampleDestination(cfg config.Destination) (core.Destination, error) {
	var dest ExampleDestination
	if err := mapstructure.Decode(cfg.Extras, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

func (d *ExampleDestination) UpdateUsers(users []string) error {
	return nil
}
