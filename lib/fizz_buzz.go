package lib

import "fmt"

func FizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzz2(num int) {
	for i := 1; i < num; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
