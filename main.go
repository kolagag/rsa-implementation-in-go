package main

import (
	"fmt"
	myRsa "myRsa/my-rsa"
)

func main() {
	pub, pri := myRsa.EncryptingKeys()

	myRsa.SaveKeysToFile("public-key.txt", pub)
	myRsa.SaveKeysToFile("private-key.txt", pri)
	fmt.Println("files are saved to current folder")

}
