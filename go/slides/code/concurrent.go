package main

import (
	"fmt"
	"time"
)

var counter = 0
// START OMIT
func main() {
	for i := 0; i < 20; i++ {
		go incr() // HL
	}

	time.Sleep(time.Millisecond * 10)
}

func incr() {
	counter++
	fmt.Println(counter)
}
// END OMIT
