package main

import (
	"flag"
	"fmt"
)

func main() {
	configFile := flag.String("c", "configFile.toml", "-c configFile file")
	fmt.Println(configFile)

	fmt.Println("Hello World")
}
