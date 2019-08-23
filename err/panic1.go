package main

import "fmt"

func raisePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in raise Panic", r)
		}
	}()

	panic("Something went wrong")
}

func main() {
	raisePanic()
	fmt.Println("Returned normally from f.")
}
