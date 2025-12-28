package myRsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

type Key struct {
	N   *big.Int
	Exp *big.Int
}

func generatePrimes(bits int) *big.Int {
	p, _ := rand.Prime(rand.Reader, bits)
	return p
}

func EncryptingKeys() (public, private Key) {

	bits := 1024
	if len(os.Args) >= 2 {
		bits, _ = strconv.Atoi(os.Args[1])
		fmt.Println("Using key size:", bits)
	} else {
		fmt.Println("No size passed -> defaulting to 1024 bits")
	}

	p := generatePrimes(bits)
	q := generatePrimes(bits)

	var d_exponent *big.Int

	n := new(big.Int).Mul(p, q)
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	q1 := new(big.Int).Sub(q, big.NewInt(1))
	eulers_quotient := new(big.Int).Mul(p1, q1)
	e_exponent := big.NewInt(65537)

	d_exponent = moduloInverse(e_exponent, eulers_quotient)
	public = Key{N: n, Exp: e_exponent}
	private = Key{N: n, Exp: d_exponent}

	return public, private

}

func moduloInverse(e, phi *big.Int) *big.Int {
	x, _, g := extendedGCD(e, phi)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil
	}

	result := new(big.Int).Mod(x, phi)
	return result
}

func extendedGCD(a, b *big.Int) (x, y, g *big.Int) {

	if b.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1), big.NewInt(0), new(big.Int).Set(a)
	}

	q := new(big.Int).Div(a, b)
	r := new(big.Int).Mod(a, b)

	x1, y1, g := extendedGCD(b, r)

	x = y1
	y = new(big.Int).Sub(x1, new(big.Int).Mul(q, y1))
	return
}

func SaveKeysToFile(filename string, key Key) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(key.Exp.String() + "\n")
	file.WriteString(key.N.String())

}
