package main

import (
	"fmt"

	"github.com/klausmark/go-public/classichelloworldmessage"
	"github.com/klausmark/go-public/stupidmessage"
)

func main() {
	var sm stupidmessage.StupidMessage = classichelloworldmessage.ClassicHelloWorldMessage{}
	message := sm.GetMessage()
	fmt.Printf("%s\n", message)
}
