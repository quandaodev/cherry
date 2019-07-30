package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
	"github.com/satori/go.uuid"
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

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Login handler")
	utils.GenerateHTML(writer, nil, "login.layout", "login")
}

func HandleAuthenticate(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Authenticate handler")
	err := request.ParseForm()
	if err != nil {
		utils.LogError("Cannot parse form", err)
	}

	username := request.PostFormValue("username")
	password := request.PostFormValue("password")

	if username != "quan" || password != "quan123" {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	v4, _ := uuid.NewV4()
	sessionToken := v4.String()

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}
