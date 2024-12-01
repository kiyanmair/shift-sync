package core

type Integration interface {
	Valid() (bool, error)
}

type Source interface {
	GetUsers() ([]string, error)
}

type Destination interface {
	SetUsers(users []string) error
}
