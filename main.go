package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{First: "Wilson"}
	p2 := person{First: "John"}
	p3 := person{First: "Jane"}

	xp := []person{p1, p2, p3}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

	xp2 := []person{}
	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(xp2)

}
