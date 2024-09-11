package router

import (
	"github.com/gorilla/mux"
	"github.com/the-jey/gomushroomapi/controllers"
	"github.com/the-jey/gomushroomapi/middlewares"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Hello, World route!
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// Mushrooms routes
	r.HandleFunc("/mushrooms", controllers.GetAllMushrooms).Methods("GET")
	r.HandleFunc("/mushrooms", controllers.DeleteAllMushrooms).Methods("DELETE")
	r.HandleFunc("/mushroom", controllers.CreateMushroom).Methods("POST")
	r.HandleFunc("/mushroom/{id}", controllers.GetOneMushroomByID).Methods("GET")
	r.HandleFunc("/mushroom/{id}", controllers.UpdateMushroomByID).Methods("PUT")
	r.HandleFunc("/mushroom/{id}", controllers.DeleteOneMushroomByID).Methods("DELETE")

	// Auth routes
	r.HandleFunc("/user/new", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")

	// Users routes
	r.HandleFunc("/users", middlewares.IsAdmin(controllers.GetAllUsers)).Methods("GET")
	r.HandleFunc("/users", middlewares.IsAdmin(controllers.DeleteAllUsers)).Methods("DELETE")
	r.HandleFunc("/user/{id}", middlewares.IsAdmin(controllers.GetUserByID)).Methods("GET")
	r.HandleFunc("/user/{id}", middlewares.IsAdmin(controllers.DeleteUserByID)).Methods("DELETE")
	r.HandleFunc("/user/username/{username}", middlewares.IsAdmin(controllers.GetUserByUsername)).Methods("GET")
	r.HandleFunc("/user/email/{email}", middlewares.IsAdmin(controllers.GetUserByEmail)).Methods("GET")

	return r
}
