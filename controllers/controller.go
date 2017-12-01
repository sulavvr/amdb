package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	// "net/http"

	"github.com/sulavvr/neogo/database"
)

type Controller interface {
}

var (
	Templates *template.Template
	DB        *sql.DB
	data      map[string]interface{}
)

func init() {
	Templates = template.Must(template.ParseGlob("views/*"))
	database.Setup()

	DB = database.Get()
	data = make(map[string]interface{})
}

func paginate(start, end, total, limit int, is string) []template.HTML {
	var pagination []template.HTML
	pagination = append(pagination, template.HTML(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"/%s\" aria-label=\"Previous\"><span aria-hidden=\"true\">&laquo;</span><span class=\"sr-only\">Previous</span></a></li>", is)))
	for i := start; i < end; i++ {
		s := template.HTML(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"/%s?skip=%d\">%d</a></li>", is, i*limit, i))
		pagination = append(pagination, s)
	}
	pagination = append(pagination, template.HTML(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"/%s?skip=%d\" aria-label=\"Previous\"><span aria-hidden=\"true\">&raquo;</span><span class=\"sr-only\">Next</span></a></li>", is, total-limit)))

	return pagination
}
