package controller

import (
	"fmt"
	"net/http"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

// HandleError shows the error message page
func HandleError(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("err - Not implemented")
	/*
		vals := request.URL.Query()
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
		}*/
}

// HandleIndex shows the index page
func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	//posts, err := model.ListPosts()
	articles, err := model.ListArticles()
	fmt.Println("Size of ariticles ", len(articles))
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get posts")
	} else {
		//_, err := session(writer, request)
		//if err != nil {
		utils.GenerateHTML(writer, articles, "layout", "public.navbar", "index")
		//} else {
		//	generateHTML(writer, posts, "layout", "private.navbar", "index")
		//}
	}
}
