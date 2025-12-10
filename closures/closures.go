package main

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100
	debitFunc := func(debitAmout int) int {
		amount -= debitAmout
		return amount
	}
	return debitFunc
}
func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))
}

/*
	debitFunc function depends on variable outside of itself, under the hood the programming language will enclose both debitFunc and the variable amount in itself and store it in memory as an enclosed unit or closure.

	The function call to activateGiftCard() treats this as seperate closure.
	-----------------------------------------------
	|	amount := 100                             |
	|	debitFunc := func(debitAmout int) int {   |
	|		amount -= debitAmout                  |
	|		return amount                         |
	|	}                                         |
	|	return debitFunc                          |
	-----------------------------------------------
*/