package services

import (
	"html/template"
	"net/http"

	"github.com/sundayfun/go-web/logs"
)

// 渲染页面并输出
func gotoHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, err := template.ParseFiles("../html/" + file + ".html")
	logs.CheckErr(err)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}
