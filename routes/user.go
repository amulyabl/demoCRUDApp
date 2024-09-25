package routes

import (
	"fmt"
	"github.com/amulya.bl/demoCRUDApp/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for:", r.URL.Path)
	http.ServeFile(w, r, "index.html")
}
func RegisterRoutes(r *mux.Router, userHandler *handlers.UserHandler) {
	r.HandleFunc("/", serveHTML).Methods(http.MethodGet)
	r.HandleFunc("/user", userHandler.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", userHandler.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.UpdateUserHandler).Methods(http.MethodPatch)
	r.HandleFunc("/user/{id}", userHandler.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/users", userHandler.GetAllUserHandler).Methods(http.MethodGet)
	//r.HandleFunc("/user/", userHandler.DeleteUserHandler).Methods(http.MethodGet)
}
