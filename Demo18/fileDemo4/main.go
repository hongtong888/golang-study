package main

import (
	"fmt"
	"io/ioutil"
)

/**
ioutils 实现文件的读取和复制
*/
func main() {
	src := "E:\\\\pom.xml"
	dst := "E:\\\\pom1.xml"
	copy(src, dst)
}

func copy(srcFileName string, descFileName string) {
	file, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err1 := ioutil.WriteFile(descFileName, file, 06666)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("复制成功")
}
