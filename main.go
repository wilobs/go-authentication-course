package main

import (
	"encoding/base64"
	"fmt"
)

type person struct {
	First string
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

}
