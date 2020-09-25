package utils

import (
	"fmt"
	"strconv"
)

func GetInt(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("utils -> GentInt : error, ", err)
		fmt.Println("Can't convert into Integer")
		fmt.Println("Please re-check .env, you recently input")
		fmt.Println(x)
	}

	return i
}
