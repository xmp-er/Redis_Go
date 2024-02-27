package process

import "strconv"

func crud(str []string) { //expectation is that the incoming command must be SET,GET, OR DEL
	if str[0] == "SET" {
		op_set(str)
	} else if str[0] == "GET" {
		op_get(str)
	} else if str[0] == "DEL" {
		op_del(str)
	}
}

func op_set(str []string) {
	k := str[1]
	v := str[2]
	temp_v, err := strconv.Atoi(v)
	if err != nil {
		//it means we have a string and we can go ahead with stripping the quotes if any and go ahead and put it in string map
		map[k]
	}
}
