package main

import (
	"fmt"
	"wulongcha/snowflake"
)

func main() {
	h := snowflake.GenInt64String()
	fmt.Println("来喝乌龙茶", h)
}
