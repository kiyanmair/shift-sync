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
	default:
		return nil, fmt.Errorf("unsupported source type: %s", config.Type)
	}
}
