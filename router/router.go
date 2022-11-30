package router

import (
	"es/controller"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	// ////////////////////////////////////////User Details////////////////////////////////////////////////
	router.HandleFunc("/api/add/user-details/", controller.InsertUserDetails)
	router.HandleFunc("/api/search/user-details/", controller.SearchByID)

	return router
}
