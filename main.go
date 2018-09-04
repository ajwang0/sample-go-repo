package main 

import (
	"fmt"
	"sample-go-repo/hello"
	"sample-go-repo/world"
)

func main() {
	fmt.Println("Hello World")
	world.World()
	hello.Hello()
}
