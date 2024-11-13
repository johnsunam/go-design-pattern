package main

import "fmt"

// embed the publication
type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is newspaper name %s", n.name)
}

func createNewsPaper(name, publisher string, page int) IPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			publisher: publisher,
			page:      page,
		},
	}
}
