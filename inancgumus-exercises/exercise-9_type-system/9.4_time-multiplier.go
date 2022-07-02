package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t, _ := time.ParseDuration("1h30m")
	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t *= time.Duration(multiplier)
	fmt.Println(t)
}
