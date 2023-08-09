package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	configFile := flag.String("c", "config_dev.toml", "-c configFile file")
	flag.Parse()

	fmt.Println("Hello World1")
	fmt.Println(*configFile)

	fmt.Println("Hello World1")
	time.Sleep(time.Hour)
}
