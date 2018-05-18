package services

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"github.com/sundayfun/go-web/db/model"
	"github.com/sundayfun/go-web/tool"
)

type Err struct {
	Error string
}

func SaveCompany(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../html/savecompany.html")
		t.Execute(w, tool.NowDate())
	} else {
		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
		money, err := strconv.ParseFloat(r.Form["money"][0], 64)
		if err != nil {
			// &Err{r.Form["money"][0]}
			gotoHTML(w, "err", r.Form["money"])
			return
		}
		gotoHTML(w, "ok", nil)
		d := &model.Company{}
		d.User = sql.NullString{r.Form["user"][0], true}
		d.Money = sql.NullFloat64{money, true}

		model.SaveCompany(d)

		// http.Redirect(w, r, "/savecompany", http.StatusFound)
		// d.Delete(db.GlobalDB)
	}
}
