package temp

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var Temp *template.Template

func InitTemplates() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
	}

	temp, err := template.New("").Funcs(funcMap).ParseGlob("./assets/Templates/*.html")
	if err != nil {
		log.Printf("Failed to parse templates: %v", err)
		fmt.Printf("Template directory contents:\n")
		files, _ := os.ReadDir("./assets/Templates")
		for _, file := range files {
			fmt.Printf("- %s\n", file.Name())
		}
		log.Fatalf("Fatal error parsing templates: %v", err)
	}

	Temp = temp
	log.Printf("Templates initialized successfully")
}
