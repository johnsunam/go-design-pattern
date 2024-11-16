package main

import "fmt"

func main() {
	fst := &first{}
	sd := &second{}

	fst.setName("compatibleName")
	fst.setAddress("compatibleAddress")
	adpt := adapter{
		scd: sd,
	}
	adpt.setName("adaptedName")
	adpt.setAddress("adaptedAddress")

	fmt.Printf("\n first %+v \n", fst)
	fmt.Printf("\n second %+v \n", sd)
}
