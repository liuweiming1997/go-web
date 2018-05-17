package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"

// 	"github.com/sundayfun/go-web/db/model"
// 	"github.com/sundayfun/go-web/logs"
// 	"github.com/sundayfun/go-web/services"
// 	"github.com/sundayfun/go-web/util"
// )

// func root(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/savecompany", http.StatusFound)
// }

// type st struct {
// 	Name []string
// }

// var s *st

// func test(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Method)
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("../html/test.html")
// 		log.Println(t.Execute(w, s))
// 	} else {
// 		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
// 	}
// }
// func init() {
// 	fmt.Println("begin")

// 	logs.LogToFile("a new start")
// 	model.XOLog = func(s string, args ...interface{}) {
// 		res := fmt.Sprintf("%s ---- %v\n", s, args)
// 		logs.LogToFile(res)
// 	}
// 	fmt.Println(util.Command("pwd"))

// 	http.HandleFunc("/", root)                                          //设置访问的路由
// 	http.HandleFunc("/savecompany", services.SaveCompany)               //设置访问的路由
// 	http.HandleFunc("/querysumforcompany", services.QuerySumForCompany) //设置访问的路由
// 	http.HandleFunc("/querysumforself", services.QuerySumForSelf)       //设置访问的路由
// 	http.HandleFunc("/test", test)                                      //设置访问的路由

// }

// func main() {
// 	err := http.ListenAndServe(":9090", nil) //设置监听的端口
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
