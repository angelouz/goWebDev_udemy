package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}
	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
