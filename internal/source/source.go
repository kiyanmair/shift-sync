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
	case "example":
		source, err := NewExampleSource(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create example source: %w", err)
		}
		return source, nil
	default:
		return nil, fmt.Errorf("unsupported source type: %s", config.Type)
	}
}
