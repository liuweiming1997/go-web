package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sundayfun/go-web/db/model"
)

type result struct {
	Command string
	Result  float64
}

func QuerySumForCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("QuerySumForCompany", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../html/querysum.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
		res, err := model.SumCompanyByUser(r.Form["user"][0])
		if err != nil {
			gotoHTML(w, "err", r.Form["user"])
			return
		}
		gotoHTML(w, "sum", &result{"公司报销金额", res})
	}
}

func QuerySumForSelf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("QuerySumForSelf", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../html/querysum.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
		res, err := model.SumSelfByUser(r.Form["user"][0])
		if err != nil {
			gotoHTML(w, "err", r.Form["user"])
			return
		}
		gotoHTML(w, "sum", &result{"个人支出金额", res})
	}
}
