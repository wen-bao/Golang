package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDb map[string]PersonInfo
	personDb = make(map[string]PersonInfo)

	personDb["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDb["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	person, ok := personDb["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
}
