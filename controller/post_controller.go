package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// NewPost handles GET post/new to display the new post page
func NewPost(writer http.ResponseWriter, request *http.Request) {
	// TODO: check session
	utils.GenerateHTML(writer, nil, "layout", "public.navbar", "new.post")
}

// CreatePost handles POST post/create to save the new post to database
func CreatePost(writer http.ResponseWriter, request *http.Request) {
	/*sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
	*/
	err := request.ParseForm()
	if err != nil {
		utils.LogError("Cannot parse form", err)
	}

	var p model.Post
	p.Title = request.PostFormValue("title")
	p.Content = request.PostFormValue("content")

	if err = model.CreatePost(p); err != nil {
		utils.LogError("Cannot create post", err)
	}
	http.Redirect(writer, request, "/", 302)
	//	}
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
