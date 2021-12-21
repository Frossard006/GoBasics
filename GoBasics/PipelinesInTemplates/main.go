package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

//FOR TIME IN TEMPLATES

/*
var tpl *template.Template

var fm = template.FuncMap{
	"fDateMDY": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("one.gohtml"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func handle(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "one.gohtml", time.Now())
}
*/

//FOR PIPELINES IN TEMPLATES

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("two.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "two.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}
