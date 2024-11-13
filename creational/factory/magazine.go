package main

import "fmt"

// embed the publication
type magazine struct {
	publication
}

func (n magazine) String() string {
	return fmt.Sprintf("This is magazine name %s", n.name)
}

func createMagazine(name, publisher string, page int) IPublication {
	return &magazine{
		publication: publication{
			name:      name,
			publisher: publisher,
			page:      page,
		},
	}
}
