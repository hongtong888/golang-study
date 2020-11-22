package main

import (
	"fmt"
	"io"
	"os"
)

/**
一、读取文件（方法1）
		1、只读方式打开文件 file,err := os.Open()

		2、读取文件 file.Read()

		3、关闭文件流 defer file.Close()
*/
func main() {
	//1、只读方式打开文件操作
	//file, err := os.Open("F:\\code\\go\\golang-study\\Demo18\\fileDemo1\\main.go")
	////文件流必须要关闭
	//defer file.Close()
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	////对文件进行操作,打印的是文件的指针地址
	//fmt.Println(file)

	//2、读取文件的内容
	file, err := os.Open("E:\\pom.xml")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	//读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		//最后一次可能内容较少，采取返回多少 n   就读取多少
		strSlice = append(strSlice, tempSlice[:n]...)
	}
	fmt.Println(string(strSlice))
}
