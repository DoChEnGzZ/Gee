package main

import (
	"GeeProject/Gee"
	"fmt"
	"html/template"
	"log"
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
	g:=Gee.Default()
	g.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAsDate,
	})
	g.Static("/ass","./static")
	g.LoadHTMLGlob("temp/*")
	g.Get("/date", func(c *Gee.Context) {
		c.HTMl(http.StatusOK, "custom_func.tmpl", Gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})
	students:=g.Group("/students")
	students.Get("/", func(c *Gee.Context) {
		c.HTMl(http.StatusOK,"css.tmpl",nil)
	})
	students.Get("/query", func(c *Gee.Context) {
		c.HTMl(http.StatusOK,"arr.tmpl",Gee.H{
			"title":"students",
			"stuArr:":[2]*student{stu1,stu2},
		})
	})
	g.Get("/panic", func(c *Gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	//ass:=g.Group("/ass")
	//ass.Post("/*filepath", func(c *Gee.Context) {
	//	c.Json(http.StatusOK,Gee.H{
	//		"filepath":c.Param("filepath"),
	//	})
	//})
	err := g.Run(":8080")
	if err != nil {
		log.Printf("s/n",err.Error())
	}
}
