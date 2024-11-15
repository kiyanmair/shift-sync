package source

import (
	"fmt"

	"github.com/kiyanmair/shift-sync/config"
)

type Source interface {
	FetchUsers() ([]string, error)
}

func NewSource(config config.Source) (Source, error) {
	switch config.Type {
	default:
		return nil, fmt.Errorf("unsupported source type: %s", config.Type)
	}
}
