package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func arrayContainsInt( array []int, number int) bool {
	for i := range array {
    	if( array[i] == number){
			return true
		}
	}
	return false
}

func missingNumber (array []int){
	fmt.Println("Find a missing number in an array")
	for i:= 1; i <= 10; i++ {
		if( arrayContainsInt( array, i)){
			continue
		} else {
			fmt.Print("Missing integer: ")
			fmt.Println(i)
		}
	}
}

func twice(f func(int) int) func(int) int {
	return func(x int) int {
		return f(f(x))
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main(){

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select an option to run:")
		fmt.Println("1. Find a missing item in an array")
		fmt.Println("2. Use a higher order function")
		fmt.Println("exit")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		switch text {
		case "1":
			array := []int{1,2,3,4,5,6,7,8,9}
			missingNumber(array)
		case "2":
			/* mess with some higher-order functions */
			plusThree := func(i int) int {
				return i + 3
			}

			g := twice(plusThree)
			fmt.Println(squareSum(2)(6)(3))
			fmt.Println(g(7)) // 13
			fmt.Println(plusThree(7))

			//s := []string{"Foo", "Bar"}
		case "exit":
			return
		default:
			fmt.Println( text)
			fmt.Println("Invalid option selected. Please use the numeric value of the option you want to select.")
		}
	}
}