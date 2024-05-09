package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println(sum(2, 5, 3))           //output 10
	fmt.Println(sum(2, 5, 3, 5, 8, 10)) //output 33
	write("Olha o número", 1, 2, 3, 4, 5, 6, 7)
	// output
	//Olha o número 1
	//Olha o número 2
	//Olha o número 3
	//Olha o número 4
	//Olha o número 5
	//Olha o número 6
	//Olha o número 7

}
