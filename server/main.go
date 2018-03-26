package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sundayfun/go-web/db/model"
	"github.com/sundayfun/go-web/logs"
	"github.com/sundayfun/go-web/services"
	"github.com/sundayfun/go-web/util"
)

func root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/savecompany", http.StatusFound)
}

func init() {
	logs.LogToFile("a new start")
	model.XOLog = func(s string, args ...interface{}) {
		res := fmt.Sprintf("%s ---- %v\n", s, args)
		logs.LogToFile(res)
	}
	fmt.Println(util.Command("pwd"))

	http.HandleFunc("/", root)                                          //设置访问的路由
	http.HandleFunc("/savecompany", services.SaveCompany)               //设置访问的路由
	http.HandleFunc("/querysumforcompany", services.QuerySumForCompany) //设置访问的路由
	http.HandleFunc("/querysumforself", services.QuerySumForSelf)       //设置访问的路由
}

func main() {
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
