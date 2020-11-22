package main

import (
	"fmt"
	"io/ioutil"
)

/**
二、读取文件（方法2）bufio 读取文件

		1、只读方式打开文件 file,err := os.Open()

		2、创建reader对象  reader := bufio.NewReader(file)

		3、ReadString读取文件  line, err := reader.ReadString('\n')

		4、关闭文件流 defer file.Close()
*/
func main() {
	//1、bufio读取文件的内容
	//file, err := os.Open("E:\\pom.xml")
	//defer file.Close()
	//if err != nil {
	//	fmt.Println("打开文件失败")
	//	return
	//}
	////bufio读取文件
	//var fileStr string
	//reader := bufio.NewReader(file)
	//for{
	//	readString, err := reader.ReadString('\n')
	//	if err == io.EOF{
	//		fmt.Println("文件读取完毕")
	//		fileStr += readString
	//		break
	//	}
	//    if err != nil{
	//    	fmt.Println("文件读取错误")
	//		return
	//	}
	//    fileStr += readString
	//}
	//fmt.Println(fileStr)

	//2、ioutil按行读取全部文件
	file, err := ioutil.ReadFile("E:\\pom.xml")
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}
	fmt.Println(string(file))
}
