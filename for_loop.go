package main

import "fmt"

func main() {

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println("Iteration:", i)
	// }

	//iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// 	fmt.Println("Index: ", index, " Value: ", value)
	// 	// %d - brojevi; %v - generalna vrednost
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Neparan broj ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Broj je ", i)
	// }

	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	// unutrasnji loop za space pre zvezdica
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// unutrasnji za zvezdice
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10 {
		fmt.Println(10 - i)
	}

}
