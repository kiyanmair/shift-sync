package core

func AsSource(i Integration) (Source, bool) {
	s, ok := i.(Source)
	return s, ok
}

func AsDestination(i Integration) (Destination, bool) {
	d, ok := i.(Destination)
	return d, ok
}

func IsSource(i Integration) bool {
	_, ok := AsSource(i)
	return ok
}

func IsDestination(i Integration) bool {
	_, ok := AsDestination(i)
	return ok
}
