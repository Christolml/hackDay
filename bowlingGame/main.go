package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
//
//
//
// 	for ronda := 0; ronda < 9; ronda++  {
// 		for turno := 0; turno < 2; turno++  {
// 			rand.Seed(time.Now().UnixNano())
// 			fmt.Println(randomInt(1, 11)) //get an int in the 1...10 range
// 		}
// 		fmt.Println("holi")
// 	}
// }
//
// func randomInt(min, max int) int {
// 	num := min + rand.Intn(max-min)
// 	return num
// }


func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomArray(10))
}