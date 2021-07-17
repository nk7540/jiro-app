package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const DELIMITER = ","

func IntSliceToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", DELIMITER, -1), "[]")
}

func StringToIntSlice(a string) []int {
	strs := strings.Split(a, DELIMITER)
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
