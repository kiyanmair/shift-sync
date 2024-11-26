package destination

type Destination interface {
	UpdateUsers(users []string) error
}
