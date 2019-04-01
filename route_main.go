package main

import (
	"net/http"

	"github.com/quandaodev/cherry/data"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	p("err - Not implemented")
	/*
		vals := request.URL.Query()
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
		}*/
}

func index(writer http.ResponseWriter, request *http.Request) {
	posts, err := data.Posts()
	if err != nil {
		error_message(writer, request, "Cannot get posts")
	} else {
		//_, err := session(writer, request)
		//if err != nil {
		generateHTML(writer, posts, "layout", "public.navbar", "index")
		//} else {
		//	generateHTML(writer, posts, "layout", "private.navbar", "index")
		//}
	}
}