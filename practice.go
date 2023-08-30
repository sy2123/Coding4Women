package main

import"fmt"

func practice() {
	myNum := []int{2, 4, 6, 8, 10}
	sum := 0
	for i:= 0 ; i < len(myNum) ; i++ {
		sum = sum + myNum[i]
	}
	fmt.Println(myNum)
	fmt.Println(sum)
}