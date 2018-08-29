package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World, env BEAUTIFUL=%s\n", os.Getenv("BEAUTIFUL"))
}
