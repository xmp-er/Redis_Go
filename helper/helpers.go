package helper

import "strconv"

func Set_db(st []string) int {
	ret, _ := strconv.Atoi(st[1])
	return ret
}
