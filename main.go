package main

import (
	"fmt"

	"nipun.io/message_queue/appcontext"
)

func main() {
	appcontext.Init()
	fmt.Println("Hello World!!")
}
