package main

import (
	"fmt"
	"time"
)

func main() {
	// after done, NewTimer would be set to channel property C,
	var t = time.NewTimer(2 * time.Second)

	fmt.Println("start newtimer")

	// we need to call proerty C for representinng t is done
	<-t.C

	fmt.Println("stop newtimer after 2 seconds")

}
