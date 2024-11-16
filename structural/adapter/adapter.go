package main

type adapter struct {
	scd *second
}

func (a *adapter) setName(name string) {
	a.scd.setTitle(name)
}

func (a *adapter) setAddress(address string) {
	a.scd.setLocation(address)
}
