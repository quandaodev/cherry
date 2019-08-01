package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// EditPost handles GET post/new to display the edit post page
func EditPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("EditPost() called")
	params := request.URL.Query()
	postID := params.Get("id")
	article, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get post")
	} else if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, article, "layout", "private.navbar", "edit.post")
	}
}

// UpdatePost handles POST update the post to database
func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("UpdatePost() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PostDB
		p.Title = request.PostFormValue("title")
		p.Slug = request.PostFormValue("slug")
		p.Markdown = request.PostFormValue("markdown")
		p.Content = request.PostFormValue("content")

		if err = model.UpdatePost(p); err != nil {
			utils.LogError("Cannot update post", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// NewPost handles GET post/new to display the new post page
func NewPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("NewPost() called")
	if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.post")
	}
}

// CreatePost handles POST post/create to save the new post to database
func CreatePost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("CreatePost() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PostDB
		p.Title = request.PostFormValue("title")
		p.Slug = request.PostFormValue("slug")
		p.Markdown = request.PostFormValue("markdown")
		p.Content = request.PostFormValue("content")

		if err = model.CreatePost(p); err != nil {
			utils.LogError("Cannot create post", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// ReadPost shows the details of the thread, including the posts and the form to write a post
func ReadPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("ReadPost() called")
	params := request.URL.Query()
	postID := params.Get("id")
	article, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get article")
	} else {
		if !HasSignedIn(request) {
			utils.GenerateHTML(writer, article, "layout", "public.navbar", "public.post")
		} else {
			utils.GenerateHTML(writer, article, "layout", "private.navbar", "public.post")
		}
	}
}
