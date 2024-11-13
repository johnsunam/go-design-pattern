package main

import "fmt"

// factory function that creates new object with respect to provided pubType
func newPublication(pubType string, name string, page int, publisher string) (IPublication, error) {
	if pubType == "newspaper" {
		return createNewsPaper(name, publisher, page), nil
	} else if pubType == "magazine" {
		return createMagazine(name, publisher, page), nil
	}
	return nil, fmt.Errorf("no such pubtype found")
}
