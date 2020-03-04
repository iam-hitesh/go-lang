package main

import "fmt"

func getSum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6}

	//  Like maps and slices, channels must be created before use:
	c := make(chan int)

	// First half of array
	go getSum(arr[:len(arr)/2], c)

	// Second half of array
	go getSum(arr[len(arr)/2:], c)

	x, y := <- c, <- c
	fmt.Println("Sum of first half: ", x)
	fmt.Println("Sum of second half: ", y)
}
