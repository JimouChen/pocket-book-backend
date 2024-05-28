package comm

import "strconv"

func Str2Int(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
