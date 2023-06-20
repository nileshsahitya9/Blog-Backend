package routes

import (
	"net/http"

	"github.com/nileshsahitya9/Blogs-Backend/internal/controllers"
)

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", controllers.HomeHandler)
	// Add more routes here

	return router
}
