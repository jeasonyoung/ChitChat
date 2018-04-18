package main

import (
	"net/http"
	"html/template"
	"github.com/jeasonyoung/ChitChat/src/data"
)

//
func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)

	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)

	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:"0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}

//index
//func index(w http.ResponseWriter, r *http.Request) {
//	files := []string{
//		"templates/layout.html",
//		"templates/navbar.html",
//		"templates/index.html",}
//
//	templates := template.Must(template.ParseFiles(files...))
//	if threads, err := data.Threads(); err == nil {
//		templates.ExecuteTemplate(w, "layout", threads)
//	}
//}