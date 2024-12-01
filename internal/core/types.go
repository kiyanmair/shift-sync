package core

type Integration interface {
	Validate(direction IntegrationDirection) error
}

type Source interface {
	Integration
	GetUsers() ([]string, error)
}

type Destination interface {
	Integration
	SetUsers(users []string) error
}

type IntegrationDirection string

const (
	SourceDirection      IntegrationDirection = "source"
	DestinationDirection IntegrationDirection = "destination"
)
