package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/quandaodev/cherry/model"
	"github.com/quandaodev/cherry/utils"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte(utils.Config.SessionStore)
	// Store is the session store
	Store = sessions.NewCookieStore(key)
)

// HandleError shows the error message page
func HandleError(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handle error")
	vals := request.URL.Query()
	session, _ := Store.Get(request, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// HandleIndex shows the index page
func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	articles, err := model.ListPosts()
	fmt.Println("Size of ariticles ", len(articles))
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get posts")
	} else {
		session, _ := Store.Get(request, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			utils.GenerateHTML(writer, articles, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(writer, articles, "layout", "private.navbar", "index")
		}
	}
}

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Login handler")
	utils.GenerateHTML(writer, nil, "login.layout", "login")
}

func HandleLogout(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Login handler")
	session, _ := Store.Get(request, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(request, writer)

	http.Redirect(writer, request, "/", 302)
}

func HandleAuthenticate(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Authenticate handler")

	err := request.ParseForm()
	if err != nil {
		utils.LogError("Cannot parse form", err)
	}

	username := request.PostFormValue("username")
	b := md5.Sum([]byte(request.PostFormValue("password")))
	password := hex.EncodeToString(b[:])
	fmt.Println(password)

	if username != utils.Config.Username || password != utils.Config.Password {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	session, _ := Store.Get(request, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(request, writer)

	http.Redirect(writer, request, "/", 302)
}
