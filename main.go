package main

import (
	"log"

	"./lib"
)

func main() {
	a := lib.SumNumList([]int{1, 2, 4, 5})
	log.Println(a)
}
