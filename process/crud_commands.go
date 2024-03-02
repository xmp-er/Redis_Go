package main

import (
	"strings"
)

//1 stands for successful op, 0/nil for failed op

func crud(str []string, Map map[string]string, Backup_Map map[string]string) string { //expectation is that the incoming command must be SET,GET, OR DEL
	var res string = ""
	if str[0] == "SET" {
		res = op_set(str, Map, Backup_Map)
	} else if str[0] == "GET" {
		res = op_get(str, Map, Backup_Map)
	} else if str[0] == "DEL" {
		res = op_del(str, Map, Backup_Map)
	}
	return res
}

func op_set(str []string, Map map[string]string, Backup_Map map[string]string) string { //SET operation implementation, assumption all input correct
	k := str[1]
	var v string = ""
	for ind, i := range str[2:] { //logic to remove the quotes and retain the spaces
		v += (strings.Trim(i, "\"'"))
		if ind != (len(str) - 1 - 2) {
			v += " "
		}
	}
	Map[k] = v
	Backup_Map[k] = v
	return "OK"
}

func op_get(str []string, Map map[string]string, Backup_Map map[string]string) string { //GET implementation, assuming the input to be correct
	k := str[1]
	if _, ok := Map[k]; ok {
		return `"` + Map[k] + `"`
	}
	return "nil"
}

func op_del(str []string, Map map[string]string, Backup_Map map[string]string) string { //DEL implementation, assuming input to be correct
	k := str[1]
	if _, ok := Map[k]; ok {
		delete(Map, k)
		delete(Backup_Map, k)
		return "(integer) 1"
	}
	return "(integer) 0"
}
