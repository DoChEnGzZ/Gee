package main

import (
	"GeeProject/Gee"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}


func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	g:=Gee.New()
	g.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAsDate,
	})
	g.Use(Gee.Logger())
	g.Static("/ass","./static")
	g.LoadHTMLGlob("temp/*")
	g.Get("/date", func(c *Gee.Context) {
		c.HTMl(http.StatusOK, "custom_func.tmpl", Gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})
	hello:=g.Group("/students")
	hello.Get("/", func(c *Gee.Context) {
		c.HTMl(http.StatusOK,"css.tmpl",nil)
	})
	hello.Get("/query", func(c *Gee.Context) {
		c.HTMl(http.StatusOK,"arr.tmpl",Gee.H{
			"title":"students",
			"stuArr:":[2]*student{stu1,stu2},
		})
	})

	//ass:=g.Group("/ass")
	//ass.Post("/*filepath", func(c *Gee.Context) {
	//	c.Json(http.StatusOK,Gee.H{
	//		"filepath":c.Param("filepath"),
	//	})
	//})
	err := g.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
