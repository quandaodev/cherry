package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

func createPost(writer http.ResponseWriter, request *http.Request) {
	utils.P("createPost - Not implemented")
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

// ReadPost shows the details of the thread, including the posts and the form to write a post
func ReadPost(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	postID := vals.Get("id")
	post, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot read post")
	} else {
		//_, err := session(writer, request)
		//if err != nil {
		utils.GenerateHTML(writer, &post, "layout", "public.navbar", "public.post")
		//} else {
		//generateHTML(writer, &post, "layout", "private.navbar", "private.post")
	}
}
