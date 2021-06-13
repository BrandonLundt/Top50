package main
import "fmt"

func arrayContainsInt( array []int, number int) bool {
	for i := 0; i < len(array); i++ {
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

		} else {
			fmt.Print("Missing integer: ")
			fmt.Println(i)
		}
	}
}

func main(){
	array := []int{1,2,3,4,5,6,7,8,9}
	missingNumber(array)
}