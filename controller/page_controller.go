package controller

import (
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// HandleEditPage handles GET post/new to display the edit post page
func HandleEditPage(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleEditPage() called")
	params := request.URL.Query()
	pageID := params.Get("id")
	article, err := model.GetPageByID(pageID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get page")
	} else if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, article, "layout", "private.navbar", "edit.page")
	}
}

// HandleUpdatePage handles POST update the post to database
func HandleUpdatePage(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleUpdatePage() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PageDB
		p.HTML = request.PostFormValue("html")
		p.Content = request.PostFormValue("content")

		if err = model.UpdatePage(p); err != nil {
			utils.LogError("Cannot update page", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// HandleNewPage handles GET post/new to display the new post page
func HandleNewPage(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleNewPage() called")
	if CheckAndSignIn(writer, request) {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.page")
	}
}

// HandleCreatePage handles POST post/create to save the new post to database
func HandleCreatePage(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleCreatePage() called")
	if CheckAndSignIn(writer, request) {
		err := request.ParseForm()
		if err != nil {
			utils.LogError("Cannot parse form", err)
		}

		var p model.PageDB
		p.HTML = request.PostFormValue("html")
		p.Content = request.PostFormValue("content")

		if err = model.CreatePage(p); err != nil {
			utils.LogError("Cannot create page", err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// HandleReadPage shows the page
func HandleReadPage(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleReadPage() called")
	params := request.URL.Query()
	pageID := params.Get("id")
	article, err := model.GetPageByID(pageID)
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get page")
	} else {
		if !HasSignedIn(request) {
			utils.GenerateHTML(writer, article, "layout", "public.navbar", "public.page")
		} else {
			utils.GenerateHTML(writer, article, "layout", "private.navbar", "public.page")
		}
	}
}
