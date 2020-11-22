package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

/**
一、写入文件（方法1）
	1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

	2、写入文件
		file.Write([]byte(str))        //写入字节切片数据
		file.WriteString("直接写入的字符串数据") //直接写入字符串数据

	3、关闭文件流 file.Close()

二、写入文件（方法2） bufio 写入文件

	1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

	2、创建writer对象  writer := bufio.NewWriter(file)

	3、将数据先写入缓存  writer.WriteString("你好golang\r\n")

	4、将缓存中的内容写入文件	writer.Flush()

	5、关闭文件流 file.Close()
*/
func main() {
	//1、正常的写入文件
	//file, err := os.OpenFile("E:\\pom1.xml", os.O_CREATE|os.O_RDWR, 0666)
	//defer file.Close()
	//if err!=nil{
	//    fmt.Println(err)
	//	return
	//}
	//var str = "这是测试数据"
	//file.Write([]byte(str))

	//2、bufio进行写入文件
	//file, err := os.OpenFile("E:\\pom1.xml", os.O_CREATE|os.O_RDWR, 0666)
	//defer file.Close()
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//writer := bufio.NewWriter(file)
	//for i := 0; i < 10; i++ {
	//	writer.WriteString("这是一条测试数据"+ strconv.Itoa(i)+"\r\n")
	//}
	//writer.Flush()

	//3、ioutil进行文件的写入
	var fileStr string
	for i := 0; i < 10; i++ {
		fileStr += "这是一条测试数据" + strconv.Itoa(i) + "\r\n"
	}
	err := ioutil.WriteFile("E:\\pom1.xml", []byte(fileStr), 06666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
