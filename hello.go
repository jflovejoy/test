package main



import "fmt"



func main() {
		var age int = 35
    fmt.Printf("hello world, you are %v years old.\n", age)
		fmt.Printf("You're age doubled is %v years old! \n", doubleYourAge(age))
		fmt.Printf("Your fortune is: ")
		fmt.Println(fortuneCookie(age))
}

func doubleYourAge(age int) int{
	return age * 2
}

func fortuneCookie(age int) string{
var fortune string
switch age {
case 35:
	fortune = "You will inherit a large cream cheese factory"
case 45:
	fortune = "You will have a surprise event"
case 55:
	fortune = "You will all your base belongs to us"
default:
	fortune = "You will have much good luck"
}

return fortune
}
