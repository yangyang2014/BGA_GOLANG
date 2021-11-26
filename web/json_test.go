package web

import (
	"encoding/json"
	"fmt"
	"testing"
)

//测试json 和结构体的转换

type Person struct {
	//ATTENTION: 首字母一定要大写
	//ATTENTION:转为json，更改别名使用`json:"xxxx"`
	Name string `json:"username"`
	Age int
	Sex string
	//ATTENTION:为空就不参与转为字符串，使用`json:",omitempty"`
	Hobby string`json:",omitempty"`
}

func TestStructToJson(t *testing.T) {
	p := Person{
		Name: "wyy",
		Age: 1,
		Sex:"man",
		Hobby: "",
	}
	res, err :=	json.Marshal(p)
	if err == nil {
		fmt.Println(string(res))
	}else {
		fmt.Println(err)
	}

}