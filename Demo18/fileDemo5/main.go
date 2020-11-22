package main

import (
	"fmt"
	"os"
)

/**
文件操作的一些其他api  文件重命名  创建文件夹  删除文件等
*/

func main() {
	//重命名
	//err := os.Rename("E:/pom1.xml","E:/pom2.xml")
	//if err!=nil {
	//	fmt.Println(err)
	//}

	//创建文件夹
	//err := os.Mkdir("E:/demo", 06666)
	//if err !=nil {
	//    fmt.Println(err)
	//}

	//创建多级目录
	//err := os.MkdirAll("E:/demo/demo1/demo2", 06666)
	//if err !=nil {
	//   fmt.Println(err)
	//}

	// /删除一个文件
	//err := os.Remove("E:/pom2.xml")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//删除一个目录
	//err := os.Remove("E:/demo/demo1/demo2")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 一次删除多个文件
	err := os.RemoveAll("E:/demo")
	if err != nil {
		fmt.Println(err)
	}
}
