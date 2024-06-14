package main

import (
	// "context"
	// "encoding/json"
	"log"
	"net/http"

	// "os"
	// "os/signal"
	// "strings"
	"time"

	// "github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// instances of renderer and mgo
var rnd *renderer.Render
var db *mgo.Database

// const to use throughout the program
const (
	hostName       string = "localhost:27017"
	dbName         string = "go_todo_go"
	collectionName string = "todo"
	port           string = ":9000"
)

// database model

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}

	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

// function to initialize the renderer
func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

// function to handle homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}
