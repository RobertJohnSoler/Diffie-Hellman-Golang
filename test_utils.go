package main

import (
	"Diffie-Hellman-Golang/utils"
	"fmt"
	"math/big"
)

func test_random_BigInts() {
	random1 := utils.GenRandBigInt1(128)
	random2 := utils.GenRandBigInt2(128)
	fmt.Println("First random is ", random1.String())
	fmt.Println("Second random is ", random2.String())
}

func test_primeCheck() {
	for i := 0; i <= 30; i++ {
		num := big.NewInt(int64(i))
		if utils.PrimeCheck(num) {
			fmt.Println(num, " is prime.")
		} else {
			fmt.Println(num, " is not prime.")
		}
	}
}

func test_primeGen() {
	random_prime := utils.PrimeGen(128)
	fmt.Println("Random prime: ", random_prime)
}

func main() {
	test_random_BigInts()
	test_primeCheck()
	test_primeGen()
}
