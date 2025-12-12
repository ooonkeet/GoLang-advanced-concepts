package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int32{1, 2, 3, 4, 5}
	numbers3 := []int64{1, 2, 3, 4, 5}
	numbers4 := []float32{1.1, 2.1, 3.1, 4.1, 5.1}
	numbers5 := []float64{1.1, 2.1, 3.1, 4.1, 5.1}
	fmt.Println(sumNumbers(numbers1))
	fmt.Println(sumNumbers(numbers2))
	fmt.Println(sumNumbers(numbers3))
	fmt.Println(sumNumbers(numbers4))
	fmt.Println(sumNumbers(numbers5))

}

func sumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

/*

we create a type of interface named Number which would contain the types int,int32,int64,float32,float64

then we create a generic function sumNumbers and to avoid the confusion of return type which contains T as a datatype to represent the interface Number and also pass a slice of the generic

In main function we take 5 slices of different type and display their sum by using function which contain generic.

*/