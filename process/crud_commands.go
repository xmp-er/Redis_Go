package process

import (
	"strconv"
	"strings"
)

func crud(str []string) { //expectation is that the incoming command must be SET,GET, OR DEL
	if str[0] == "SET" {
		op_set(str)
	} else if str[0] == "GET" {
		op_get(str)
	} else if str[0] == "DEL" {
		op_del(str)
	}
}

func op_set(str []string) { //SET operation implementation, assumption all input correct
	k := str[1]
	v := str[2]
	temp_v, err := strconv.ParseFloat(v, 64)
	if err != nil {
		//it means we have a float64 and we can assign it to Map_int
		Map_int[k] = temp_v
		return
	}
	//now the remaining can only be the string value so adding it in Map_string
	Map_string[k] = strings.Trim(v, "\"")
}
