package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
1、ioutils 实现文件的读取和复制
2、使用普通的文件流读取和复制
*/
func main() {
	src := "E:\\\\pom.xml"
	dst := "E:\\\\pom1.xml"
	//copy(src, dst)
	copyFile(src, dst)
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

func copyFile(srcFileName string, descFileName string) (err error) {
	//读取源文件
	file, err := os.Open(srcFileName)
	defer file.Close()
	if err != nil {
		return err
	}
	//创建目标文件
	openFile, err1 := os.OpenFile(descFileName, os.O_CREATE|os.O_WRONLY, 06666)
	defer openFile.Close()
	if err1 != nil {
		return err1
	}
	//将源文件流读取到目标文件中
	var tempSlice = make([]byte, 1280)
	for {
		read, err := file.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//复制到目标文件流
		_, err = openFile.Write(tempSlice[:read])
		if err != nil {
			return err
		}
	}
	return nil
}
