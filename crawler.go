package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reQQEmail = `(\d+)@qq.com`
)

func GetEmail() {
	//去网站获取数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleErr(err, "http get url err: ")

	defer resp.Body.Close()
	//读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, "ioutil readall failed: ")
	//把字节转成字符串
	pageStr := string(pageBytes)
	//过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	//获取结果，-1代表获取到全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//遍历结果打印出来
	for _, result := range results {
		fmt.Println("email: ", result[0])
		fmt.Println("qq: ", result[1])
	}
}
func HandleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
func main() {
	//获取邮件
	GetEmail()
}
