package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sundayfun/go-web/db/model"
	"github.com/sundayfun/go-web/logs"
	"github.com/sundayfun/go-web/services"
	"github.com/sundayfun/go-web/util"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "???\n\nfuck")
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func init() {
	logs.LogToFile("a new start")
	model.XOLog = func(s string, args ...interface{}) {
		res := fmt.Sprintf("%s ---- %v\n", s, args)
		logs.LogToFile(res)
	}
	fmt.Println(util.Command("pwd"))
	http.HandleFunc("/", sayhelloName)        //设置访问的路由
	http.HandleFunc("/login", services.Login) //设置访问的路由
}

func main() {
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
