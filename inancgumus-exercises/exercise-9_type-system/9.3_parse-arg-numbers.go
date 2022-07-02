package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is :", int8(val))

	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val))

	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))

	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", val)

	val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[5], int8(val))
}
