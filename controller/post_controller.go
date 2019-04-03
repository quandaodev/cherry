package main

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
)

func createPost(writer http.ResponseWriter, request *http.Request) {
	p("createPost - Not implemented")
	/*
			sess, err := session(writer, request)
			if err != nil {
				http.Redirect(writer, request, "/login", 302)
			} else {
				err = request.ParseForm()
				if err != nil {
					danger(err, "Cannot parse form")
				}
				user, err := sess.User()
				if err != nil {
					danger(err, "Cannot get user from session")
				}
				topic := request.PostFormValue("topic")
				if _, err := user.CreateThread(topic); err != nil {
					danger(err, "Cannot create thread")
				}
				http.Redirect(writer, request, "/", 302)
		  }
	*/
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readPost(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	postID := vals.Get("id")
	post, err := model.GetPostByID(postID)
	if err != nil {
		errorMessage(writer, request, "Cannot read post")
	} else {
		//_, err := session(writer, request)
		//if err != nil {
		generateHTML(writer, &post, "layout", "public.navbar", "public.post")
		//} else {
		//generateHTML(writer, &post, "layout", "private.navbar", "private.post")
	}
}
