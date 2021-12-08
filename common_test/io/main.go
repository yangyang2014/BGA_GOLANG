package main

import (
	"fmt"
	"os"
)

func main()  {
	WriteOveride()
}

//测试 写文件，覆盖写的问题
func WriteOveride() {
	//需要使用http请求的内容，因为copy的src是reader
	//resp,err := http.Get("http://www.baidu.com")

	//resp,err := http.Get("https://file-plaso.oss-cn-hangzhou.aliyuncs.com/dev/liveclass/dev/mini/wyytest/am3u8")
	//if err!=nil {
	//	fmt.Fprint(os.Stderr,"GET URL ERR",err)
	//}
	//defer resp.Body.Close()
	localFile := "./testOverWyy.txt"
	//fd, err := os.OpenFile(os.O_CREATE|os.O_WRONLY, 0755)//覆盖写
	//**易错点 lunix 的表现为文件没有则创建，有则截断清空；window下表现为文件没有则创建，有则截断写
	//fd, err := os.OpenFile(localFile,os.O_TRUNC|os.O_CREATE, 0755)
	//lunix下表现为清空再写
	fd, err := os.OpenFile(localFile,os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("open file %v fail,%v\n", localFile, err)
		return
	}

	defer fd.Close()
	//io.Copy(fd, resp.Body)
	fd.Write([]byte("1111111111111111111"))
}

// TODO 需要熟练系统的学习下文件操作命令


// TODO 1 读目录下的文件（一级目录）

// TODO 2 将指定文件转为json

// TODO 3 判断文件是否存在