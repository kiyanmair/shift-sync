package destination

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/config"
)

type Destination interface {
	UpdateUsers(users []string) error
}

func NewDestination(config config.Destination) (Destination, error) {
	switch config.Type {
	case "example":
		dest, err := NewExampleDestination(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create example destination: %w", err)
		}
		return dest, nil
	default:
		return nil, fmt.Errorf("unsupported source type: %s", config.Type)
	}
}
