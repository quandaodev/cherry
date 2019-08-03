package main

import (
	"net/http"
	"time"

	"github.com/quandaodev/cherry/controller"
	"github.com/quandaodev/cherry/utils"
)

func version() (v string) {
	v = "0.1"
	return
}

func main() {
	utils.P("Cherry Personal Blog", version(), "started at", utils.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// handle index, error, authenticate
	mux.HandleFunc("/", controller.HandleIndex)
	mux.HandleFunc("/error", controller.HandleError)

	mux.HandleFunc(utils.Config.LoginURL, controller.HandleLogin)
	mux.HandleFunc("/authenticate", controller.HandleAuthenticate)
	mux.HandleFunc("/logout", controller.HandleLogout)

	// handle post
	mux.HandleFunc("/post", controller.HandleReadPost)
	mux.HandleFunc("/post/new", controller.HandleNewPost)       // UI
	mux.HandleFunc("/post/create", controller.HandleCreatePost) // Save
	mux.HandleFunc("/post/edit", controller.HandleEditPost)     // UI
	mux.HandleFunc("/post/update", controller.HandleUpdatePost) // Edit post

	// handle page
	mux.HandleFunc("/page", controller.HandleReadPage)

	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
