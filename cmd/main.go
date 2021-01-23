package main

import (
	"fmt"

	"github.com/codonex/rememberme/v1"
)

func main() {
	f := rememberme.RememberYou()
	result := f("Ahmet")
	fmt.Println(result)
	result1 := f("Ahmet")
	fmt.Println(result1)
	result2 := f("Yusuf")
	fmt.Println(result2)
}
