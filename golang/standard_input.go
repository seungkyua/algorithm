package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y string
	fmt.Scanf("%s %s", &x, &y)
	/* a, _ := strconv.Atoi(x)
	b, _ := strconv.Atoi(y) */
	a, _ := strconv.ParseInt(x, 0, 64)
	b, _ := strconv.ParseInt(y, 0, 64)

	fmt.Println(a + b)

}
