package main

import (
	"./addition"
	"fmt"
)

func main() {
	finish := false
	for finish == false {

		var typeVal string
		fmt.Println("Valid types are: uint32, int32, int64 and float64")
		fmt.Println("Enter the valid type to calculate with:>")
		fmt.Scan(&typeVal)
		calcType := string(typeVal)
		fmt.Println("Type to be calculated with is: ", typeVal)

		if calcType == "uint32" {
			var inputVal int
			inputVal2 := 0
			fmt.Println("Enter the first and second number. Press enter after each number")
			fmt.Scan(&inputVal)
			val1 := uint32(inputVal)
			fmt.Println("Input was read as", val1)
			fmt.Scan(&inputVal2)
			val2 := uint32(inputVal2)
			fmt.Println("Input was read as", val2, " calculating numbers")
			fmt.Println("The sum is: ", addition.SumUint32(val1, val2))
		}
		if calcType == "int32" {
			var inputVal int
			inputVal2 := 0
			fmt.Println("Enter the first and second number. Press enter after each number")
			fmt.Scan(&inputVal)
			val1 := int32(inputVal)
			fmt.Println("Input was read as", val1)
			fmt.Scan(&inputVal2)
			val2 := int32(inputVal2)
			fmt.Println("Input was read as", val2, " calculating numbers")
			fmt.Println("The sum is: ", addition.SumInt32(val1, val2))
		}
		if calcType == "int64" {
			var inputVal int
			inputVal2 := 0
			fmt.Println("Enter the first and second number. Press enter after each number")
			fmt.Scan(&inputVal)
			val1 := int64(inputVal)
			fmt.Println("Input was read as", val1)
			fmt.Scan(&inputVal2)
			val2 := int64(inputVal2)
			fmt.Println("Input was read as", val2, " calculating numbers")
			fmt.Println("The sum is: ", addition.SumInt64(val1, val2))
		}
		if calcType == "float64" {
			var inputVal int
			inputVal2 := 0
			fmt.Println("Enter the first and second number. Press enter after each number")
			fmt.Scan(&inputVal)
			val1 := float64(inputVal)
			fmt.Println("Input was read as", val1)
			fmt.Scan(&inputVal2)
			val2 := float64(inputVal2)
			fmt.Println("Input was read as", val2, " calculating numbers")
			fmt.Println("The sum is: ", addition.SumFloat64(val1, val2))
		}

		var exit string
		fmt.Println("Exit the app? Type yes to exit, no to keep caculating ")
		fmt.Scan(&exit)
		if exit == "yes" {
			finish = true
		}
		if exit == "no" {
			finish = false
		}
	}
}
