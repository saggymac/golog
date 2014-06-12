package main

import (
	"log"
	"net/http"
	"html/template"
)

func loadTemplate( templateName string ) *template.Template {
	t, _ := template.ParseFiles( "tmpl/" + templateName)
	return t
}


func rootHandler( w http.ResponseWriter, r *http.Request ) {
	tmpl := loadTemplate( "index.html")
	if tmpl != nil {
		tmpl.Execute( w, nil)
	}
}



func main() {
	http.HandleFunc( "/", rootHandler)
	log.Fatal( http.ListenAndServe( ":8080", nil))
}


