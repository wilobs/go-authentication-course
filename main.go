package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{First: "Wilson"}
	// p2 := person{First: "John"}
	// p3 := person{First: "Jane"}

	// xp := []person{p1, p2, p3}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bs))

	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Wilson",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoding bad data", err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {

}
