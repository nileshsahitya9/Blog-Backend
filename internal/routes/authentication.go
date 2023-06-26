// routes/authentication.go

package routes

import (
	"github.com/gorilla/mux"
	"github.com/nileshsahitya9/Blogs-Backend/internal/handlers"
)

func SetupAuthenticationRoutes(router *mux.Router) {
	router.Path("/register").HandlerFunc(handlers.RegisterHandler).Methods("POST")
	router.Path("/login").HandlerFunc(handlers.LoginHandler).Methods("POST")
}
