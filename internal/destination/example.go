package destination

import (
	"github.com/go-viper/mapstructure/v2"

	"github.com/kiyanmair/shift-sync/config"
)

type ExampleDestination struct {
	ID      string
	Token   string `mapstructure:"token"`
	GroupID string `mapstructure:"group_id"`
}

func NewExampleDestination(cfg config.Destination) (Destination, error) {
	var dest ExampleDestination
	dest.ID = cfg.ID

	if err := mapstructure.Decode(cfg.Details, &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}

func (d *ExampleDestination) UpdateUsers(users []string) error {
	return nil
}
