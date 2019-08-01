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
	utils.LogInfo("HandleError() called")
	vals := request.URL.Query()
	if !HasSignedIn(request) {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// HandleIndex shows the index page
func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleIndex() called")
	articles, err := model.ListPosts()
	if err != nil {
		utils.ErrorMessage(writer, request, "Cannot get posts")
	} else {
		if !HasSignedIn(request) {
			utils.GenerateHTML(writer, articles, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(writer, articles, "layout", "private.navbar", "index")
		}
	}
}

// HandleLogin shows the login page
func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleLogin() called")
	utils.GenerateHTML(writer, nil, "login.layout", "login")
}

// HandleLogout revoke the session and cookie
func HandleLogout(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleLogout() called")
	session, _ := Store.Get(request, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(request, writer)

	http.Redirect(writer, request, "/", 302)
}

// HandleAuthenticate checks the username and password
func HandleAuthenticate(writer http.ResponseWriter, request *http.Request) {
	utils.LogInfo("HandleAuthenticate() called")
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

func CheckAndSignIn(writer http.ResponseWriter, request *http.Request) (loggedIn bool) {
	loggedIn = HasSignedIn(request)
	if !loggedIn {
		http.Redirect(writer, request, utils.Config.LoginURL, 302)
	}
	return
}

func HasSignedIn(request *http.Request) (loggedIn bool) {
	loggedIn = true
	session, _ := Store.Get(request, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		loggedIn = false
	}
	return
}
