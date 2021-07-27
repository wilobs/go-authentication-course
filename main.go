package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	h := sha256.New()

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("couldnt io.copy", err)
	}

	fmt.Printf("Here's the type BEFORE Sum: %T\n", h)
	fmt.Printf("%v\n", h)
	xb := h.Sum(nil)
	fmt.Printf("Here's the type AFTER Sum: %T\n", xb)
	fmt.Printf("%x\n", xb)

	xb = h.Sum(nil)
	fmt.Printf("Here's the type AFTER SECOND Sum: %T\n", xb)
	fmt.Printf("%x\n", xb)

	xb = h.Sum(xb)
	fmt.Printf("Here's the type AFTER THIRD Sum: %T\n", xb)
	fmt.Printf("%x\n", xb)

}
