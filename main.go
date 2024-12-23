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

func main() {
	random1, _ := genRandBigInt1(128)
	random2, _ := genRandBigInt2(128)
	fmt.Println("First random is ", random1.String())
	fmt.Println("Second random is ", random2.String())
}
