package main

type first struct {
	name    string
	address string
}

func (f *first) setName(name string) {
	f.name = name
}

func (f *first) setAddress(address string) {
	f.address = address
}
