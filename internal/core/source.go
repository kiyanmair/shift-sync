package core

type Source interface {
	FetchUsers() ([]string, error)
}
