package main

import (
	"fmt"

	"github.com/odogwuVal/cryptit/decrypt"
	"github.com/odogwuVal/cryptit/encrypt"
)

func main() {
	ecryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println("encrypted string is: ", ecryptedStr)
	fmt.Println("encrypted string is: ", decrypt.Nimbus(ecryptedStr))
}
