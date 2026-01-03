package main

import (
	"flag"
	"fmt"
	myRsa "myRsa/my-rsa"
	"os"
	"strings"
)

func main() {

	bits := flag.Int("bits", 1024, "RSA KEY SIZE")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: go run . [-bits 2048] \"your message\"")
		fmt.Println("Usage: go run . <your message> \n")
		os.Exit(1)

	}

	message := strings.Join(flag.Args(), " ")

	pub, pri := myRsa.EncryptingKeys(*bits)

	myRsa.SaveKeysToFile("public-key.txt", pub)
	myRsa.SaveKeysToFile("private-key.txt", pri)
	//fmt.Println("files are saved to current folder")
	EncryptedMessage := myRsa.Encrytion(message, pub)
	fmt.Println(myRsa.Decryption(EncryptedMessage, pri))

}
