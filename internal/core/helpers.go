package core

func IsSource(i Integration) (Source, bool) {
	s, ok := i.(Source)
	return s, ok
}

func IsDestination(i Integration) (Destination, bool) {
	d, ok := i.(Destination)
	return d, ok
}
