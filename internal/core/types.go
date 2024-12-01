package core

type Integration interface {
	Valid() (bool, error)
}

type Source interface {
	Integration
	GetUsers() ([]string, error)
}

type Destination interface {
	Integration
	SetUsers(users []string) error
}
