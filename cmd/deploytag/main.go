package main

import (
	"flag"
	"fmt"
	"time"
)

var env = flag.String("env", "prod", "target environment")

func main() {
	flag.Parse()

	now := time.Now().UTC().Format("20060102t150405")
	tag := fmt.Sprintf("%v-%s", *env, now)
	fmt.Println(tag)
}
