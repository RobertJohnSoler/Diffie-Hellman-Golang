package utils

import (
	"crypto/rand"
	"math/big"
)

func GenRandBigInt1(bit_len int) *big.Int {
	byte_len := (bit_len + 7) / 8
	rand_bytes := make([]byte, byte_len)
	rand.Read(rand_bytes)
	randBigInt := new(big.Int)
	randBigInt.SetBytes(rand_bytes)
	return randBigInt
}

func GenRandBigInt2(bit_len int) *big.Int {
	x := big.NewInt(2)
	y := big.NewInt(int64(bit_len))
	max := new(big.Int)
	max.Exp(x, y, nil)
	randBigInt, _ := rand.Int(rand.Reader, max)
	return randBigInt
}

func PrimeCheck(number *big.Int) bool {
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

func PrimeGen(size int) *big.Int {
	rand_num := GenRandBigInt1(size)
	for {
		if PrimeCheck(rand_num) == true {
			break
		} else {
			rand_num = new(big.Int).Add(rand_num, big.NewInt(1))
		}
	}
	return rand_num
}
