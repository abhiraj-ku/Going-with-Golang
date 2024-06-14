package main

import (
	// "context"
	// "encoding/json"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	// "os"
	// "os/signal"
	// "strings"
	"time"

	// "github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/middleware"
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

// function to handle createTodo api
func createTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// graceFully handle the server close due to any signal(ctrl+c)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())

	// server creation
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	// channel to start the server

	go func() {
		log.Fatal("server listening on port", port)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("listen :%s\n", err)
		}

	}()

	// waits for the signal to be sent to stopChan channel if interupt is recieved then only this code will run

	<-stopChan
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()

	log.Println("Server stoppped gracefully")

}

// todo:= func  todoHandler
func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", deleteTodo)

	})
	return rg
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}
