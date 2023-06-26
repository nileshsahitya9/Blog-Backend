package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/nileshsahitya9/Blogs-Backend/internal/constants"
)

const APIBaseURL = constants.API_PREFIX_URL

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	SetupAuthenticationRoutes(router.PathPrefix(fmt.Sprintf("%s/authentication", APIBaseURL)).Subrouter())

	return router
}
