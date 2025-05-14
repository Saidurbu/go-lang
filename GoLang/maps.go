package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map

	m := make(map[string]string)

	m["name"] = "Saidur"
	m["address"] = "Tetulia"
	fmt.Println(m["name"], m["address"])

	// IMP: if key does not exists in the map then it returns zero value

	delete(m, "address")
	clear(m)

	fmt.Println(m)

	s := map[string]int{"price": 4000, "phones": 50}
	v, ok := s["phones"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m1, m2))

}
