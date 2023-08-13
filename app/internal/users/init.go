package users

import "fmt"

var EPDB = make(map[string]interface{})

func Init() map[string]interface{} {
	EPDB["list"] = List
	fmt.Println("Users initialized")
	return EPDB
}
