package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func genRandBigInt1(bit_len int) (*big.Int, error) {
	byte_len := (bit_len + 7) / 8
	rand_bytes := make([]byte, byte_len)
	_, err := rand.Read(rand_bytes)
	if err != nil {
		return nil, err
	}
	randBigInt := new(big.Int)
	randBigInt.SetBytes(rand_bytes)
	return randBigInt, nil
}

func genRandBigInt2(bit_len int) (*big.Int, error) {
	x := big.NewInt(2)
	y := big.NewInt(int64(bit_len))
	max := new(big.Int)
	max.Exp(x, y, nil)
	randBigInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}
	return randBigInt, err
}

func primeCheck(number *big.Int) bool {
	prime := false
	s := 150

	for i := 1; i <= s; i++ {
		if number.Cmp(big.NewInt(4)) == 0 {
			break
		} else if (number.Cmp(big.NewInt(4))) < 0 {
			prime = true
			break
		}

		max := new(big.Int).Sub(number, big.NewInt(2))
		Range := new(big.Int).Sub(max, big.NewInt(2))
		random := rand.Reader
		randrange, _ := rand.Int(random, Range)
		// modded_randrange := new(big.Int).Mod(randrange, Range)
		a := new(big.Int).Add(randrange, big.NewInt(2))
		primecheck := new(big.Int)
		primecheck.Exp(a, new(big.Int).Sub(number, big.NewInt(1)), number)
		if primecheck.Cmp(big.NewInt(1)) != 0 {
			prime = false
			break
		}
		prime = true

	}
	return prime
}

func main() {
	random1, _ := genRandBigInt1(128)
	random2, _ := genRandBigInt2(128)
	fmt.Println("First random is ", random1.String())
	fmt.Println("Second random is ", random2.String())
	for i := 0; i <= 30; i++ {
		num := big.NewInt(int64(i))
		fmt.Println(num, " is ", primeCheck(num))
	}

}
