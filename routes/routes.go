package routes

import (
	"fiqri_muhammad/student-api/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/api/students", controllers.CreateStudent).Methods("POST")


	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	return c.Handler(router).(*mux.Router)
}
