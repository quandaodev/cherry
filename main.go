package main

import (
	"net/http"
	"time"
)

func main() {
	p("Personal blog", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// defined in route_main.go
	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)

	// defined in route_auth.go
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)

	// defined in route_post.go
	//mux.HandleFunc("/post/create", createPost)
	mux.HandleFunc("/post", readPost)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
