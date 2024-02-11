package main

import (
	"fmt"

	"github.com/Mohammad-Hakemi22/GenericList/genericlist"
)

func main() {
	fmt.Println("Generic list structures")
	list := genericlist.New[string]()

	list.Insert("alice")
	list.Insert("bob")
	list.Insert("alex")
	fmt.Printf("%+v", list)
	s, err := list.GetByIndex(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)
}
