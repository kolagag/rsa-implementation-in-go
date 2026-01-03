package myRsa

import "math/big"

func Decryption(cipherText []*big.Int, private Key) string {
	var plain_message []byte
	for i := range cipherText {
		decrypt := new(big.Int)
		decrypt.Exp(cipherText[i], private.Exp, private.N)
		message := decrypt.Bytes()
		plain_message = append(plain_message, message...)

	}

	return string(plain_message)

}
