package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/undestanding-mongodb/controllers"
	mongo "gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mongo.Session {
	// Connect to our local mongo
	s, err := mongo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
