package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体转json就用了反射
*/
func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Println("aa")
	var s, _ = json.Marshal(s1)
	fmt.Printf("%T\n", s)
	jsonStr := string(s)
	fmt.Println(jsonStr)
}

type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}
