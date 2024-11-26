package core

type Destination interface {
	UpdateUsers(users []string) error
}
