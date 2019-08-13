package main

import (
	"fmt"

	"github.com/huynvk/gosample/fundamentals"
	"github.com/huynvk/gosample/greeting"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	greeting.Hello()
	fundamentals.Print()
}
