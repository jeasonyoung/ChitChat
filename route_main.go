package main

import (
	"github.com/jeasonyoung/ChitChat/data"
	"net/http"
)

//GET /err?msg=
//shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	if _, err := session(writer, request); err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	if threads, err := data.Threads(); err != nil {
		error_message(writer, request, "Cannot get threads:" + err.Error())
	} else {
		if _, err := session(writer, request); err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
