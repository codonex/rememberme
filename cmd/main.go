package main

import (
	"fmt"

	"github.com/codonex/rememberme"
)

func main() {
	f := rememberme.RememberMe()
	result, _ := f("Ahmet")
	fmt.Println(result)
	result1, _ := f("Ahmet")
	fmt.Println(result1)
	result2, _ := f("Yusuf")
	fmt.Println(result2)
}
