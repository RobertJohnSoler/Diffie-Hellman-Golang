package main

import (
	"fmt"
)

func main() {
	random1 := genRandBigInt1(128)
	random2 := genRandBigInt2(128)
	fmt.Println("First random is ", random1.String())
	fmt.Println("Second random is ", random2.String())
	// for i := 0; i <= 30; i++ {
	// 	num := big.NewInt(int64(i))
	// 	fmt.Println(num, " is ", primeCheck(num))
	// }
	random_prime := primeGen(128)
	fmt.Println("Random prime: ", random_prime)
}
