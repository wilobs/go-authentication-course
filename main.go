package main

import (
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

	// http.HandleFunc("/encode", foo)
	// http.ListenAndServe(":8080", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {

}

func bar(w http.ResponseWriter, r *http.Request) {

}
