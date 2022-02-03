package routes

import (
	"fmt"
	"log"
	"net/http"

	"brunomelo.crud/v1/ioc"
	"brunomelo.crud/v1/middleware"
	"github.com/gorilla/mux"
)

func RouterHandlers() {

	router := mux.NewRouter()
	go middlewareSetup(router)

	dip := ioc.NewContainerDI()

	customer := NewCustomerRoute(dip.CustomerServices)
	config := dip.Config

	router.HandleFunc("/api/customer", customer.Get).Methods("GET")
	router.HandleFunc("/api/customer", customer.Post).Methods("POST")
	router.HandleFunc("/api/customer/{id:[0-9]+}", customer.FindById).Methods("GET")
	router.HandleFunc("/api/customer/{id:[0-9]+}", customer.Update).Methods("PUT")
	router.HandleFunc("/api/customer/{id:[0-9]+}", customer.Delete).Methods("DELETE")

	fmt.Println(fmt.Sprint("Running server on the path: http://localhost:", config.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), router))
}

func middlewareSetup(r *mux.Router) {
	r.Use(middleware.ContentyTypeMiddleware)
}
