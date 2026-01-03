package myRsa

import (
	"fmt"
	"math/big"
)

func Encrytion(message string, public Key) []*big.Int {

	//var chunks []*big.Int
	var cipherTexts []*big.Int
	chunks := divide_message_to_chunks(message, public)
	for i := range chunks {
		cipherText := new(big.Int)
		cipherText.Exp(chunks[i], public.Exp, public.N)
		cipherTexts = append(cipherTexts, cipherText)

	}

	return cipherTexts

}
func divide_message_to_chunks(message string, public Key) []*big.Int {

	chunk_size := 2
	message_into_bytes := []byte(message)
	var chunks []*big.Int
	for i := 0; i < len(message_into_bytes); i += chunk_size {
		end := i + chunk_size
		if end > len(message_into_bytes) {
			end = len(message_into_bytes)
		}

		chunk_bytes := message_into_bytes[i:end]

		chunk := new(big.Int).SetBytes(chunk_bytes)
		chunks = append(chunks, chunk)
		if chunk.Cmp(public.N) > 0 {
			fmt.Println("chunk_size is ", chunk_size, "reduce it")
			break
		}

	}

	return chunks

}
