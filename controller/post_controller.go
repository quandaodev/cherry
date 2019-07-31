package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// EditPost handles GET post/new to display the edit post page
func EditPost(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	postID := params.Get("id")
	article, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get article")
	} else {
		session, _ := Store.Get(request, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(writer, request, "/login", 302)
		} else {
			utils.GenerateHTML(writer, article, "layout", "private.navbar", "edit.post")
		}
	}
}

// UpdatePost handles POST update the post to database
func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	session, _ := Store.Get(request, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	err := request.ParseForm()
	if err != nil {
		utils.LogError("Cannot parse form", err)
	}

	var a model.PostDB
	a.Title = request.PostFormValue("title")
	a.Slug = request.PostFormValue("slug")
	a.Markdown = request.PostFormValue("markdown")
	a.Content = request.PostFormValue("content")

	if err = model.UpdatePost(a); err != nil {
		utils.LogError("Cannot update article", err)
	}
	http.Redirect(writer, request, "/", 302)
}

// NewPost handles GET post/new to display the new post page
func NewPost(writer http.ResponseWriter, request *http.Request) {
	session, _ := Store.Get(request, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	utils.GenerateHTML(writer, nil, "layout", "public.navbar", "new.post")
}

// CreatePost handles POST post/create to save the new post to database
func CreatePost(writer http.ResponseWriter, request *http.Request) {
	session, _ := Store.Get(request, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	err := request.ParseForm()
	if err != nil {
		utils.LogError("Cannot parse form", err)
	}

	var a model.PostDB
	a.Title = request.PostFormValue("title")
	a.Slug = request.PostFormValue("slug")
	a.Markdown = request.PostFormValue("markdown")
	a.Content = request.PostFormValue("content")

	if err = model.CreatePost(a); err != nil {
		utils.LogError("Cannot create article", err)
	}
	http.Redirect(writer, request, "/", 302)
}

// ReadPost shows the details of the thread, including the posts and the form to write a post
func ReadPost(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	postID := params.Get("id")
	article, err := model.GetPostByID(postID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get article")
	} else {
		session, _ := Store.Get(request, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			utils.GenerateHTML(writer, article, "layout", "public.navbar", "public.post")
		} else {
			utils.GenerateHTML(writer, article, "layout", "private.navbar", "public.post")
		}
	}
}
