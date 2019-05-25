package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Configuration contains global configuration of the website
type Configuration struct {
	Address            string
	ReadTimeout        int64
	WriteTimeout       int64
	Static             string
	MongoServerAddress string
	MongoUsername      string
	MongoPassword      string
	MongoDatabaseName  string
}

// Config contains global configuration
var Config Configuration

var logger *log.Logger

// P is a convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("cherry.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// ErrorMessage is a convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

/*
// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}*/

// ParseTemplateFiles gets in a list of file names and return a template
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// GenerateHTML generates view output
func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// LogInfo is used to write INFO log
func LogInfo(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// LogError is used to write ERROR log
func LogError(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

// LogWarning is used to write WARNING log
func LogWarning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
