package core

type Integration interface {
	Validate(direction Direction) error
}

type Source interface {
	Integration
	GetUsers() ([]string, error)
}

type Destination interface {
	Integration
	SetUsers(users []string) error
}

type Direction string

const (
	SourceDirection      Direction = "source"
	DestinationDirection Direction = "destination"
)
