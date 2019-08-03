package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// HandleEditPost handles GET post/new to display the edit post page
func HandleEditPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleEditPost() called")
	params := request.URL.Query()
	postID := params.Get("id")
	post, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get post")
	} else if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, post, "layout", "private.navbar", "edit.post")
	}
}

// HandleUpdatePost handles POST update the post to database
func HandleUpdatePost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleUpdatePost() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PostDB
		p.Title = request.PostFormValue("title")
		p.Slug = request.PostFormValue("slug")
		p.Content = request.PostFormValue("content")
		p.HTML = request.PostFormValue("html")

		if err = model.UpdatePost(p); err != nil {
			utils.LogError("Cannot update post", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// HandleNewPost handles GET post/new to display the new post page
func HandleNewPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleNewPost() called")
	if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.post")
	}
}

// HandleCreatePost handles POST post/create to save the new post to database
func HandleCreatePost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleCreatePost() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PostDB
		p.Title = request.PostFormValue("title")
		p.Slug = request.PostFormValue("slug")
		p.Content = request.PostFormValue("content")
		p.HTML = request.PostFormValue("html")

		if err = model.CreatePost(p); err != nil {
			utils.LogError("Cannot create post", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// ReadPost shows the details of the thread, including the posts and the form to write a post
func HandleReadPost(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleReadPost() called")
	params := request.URL.Query()
	postID := params.Get("id")
	post, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get post")
	} else {
		if !HasSignedIn(request) {
			utils.GenerateHTML(writer, post, "layout", "public.navbar", "public.post")
		} else {
			utils.GenerateHTML(writer, post, "layout", "private.navbar", "private.post")
		}
	}
}
