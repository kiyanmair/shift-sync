package source

type Source interface {
	FetchUsers() ([]string, error)
}
