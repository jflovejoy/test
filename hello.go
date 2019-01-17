package main



import "fmt"



func main() {
		var age int = 32
    fmt.Printf("hello world, you are %v years old.\n", age)
		fmt.Printf("You're age doubled is %v years old! \n", doubleYourAge(age))
}

func doubleYourAge(age int) int{
	return age * 2
}
