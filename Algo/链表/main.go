package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(2)
	fmt.Println(r)
}
