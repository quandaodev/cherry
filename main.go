package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/quandaodev/cherry/controller"
	"github.com/quandaodev/cherry/utils"
)

func version() (v string) {
	v = "0.1"
	return
}

func main() {
	//-----------------------
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	/*statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}*/

	//-----------------------
	utils.P("Cherry Personal Blog", version(), "started at", utils.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// handle main pages
	mux.HandleFunc("/", controller.HandleIndex)
	mux.HandleFunc("/error", controller.HandleError)

	mux.HandleFunc(utils.Config.LoginURL, controller.HandleLogin)
	mux.HandleFunc("/authenticate", controller.HandleAuthenticate)
	mux.HandleFunc("/logout", controller.HandleLogout)

	// handle post
	mux.HandleFunc("/post", controller.ReadPost)
	mux.HandleFunc("/post/new", controller.NewPost)       // UI
	mux.HandleFunc("/post/create", controller.CreatePost) // Save
	mux.HandleFunc("/post/edit", controller.EditPost)     // UI
	mux.HandleFunc("/post/update", controller.UpdatePost) // Edit post

	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
