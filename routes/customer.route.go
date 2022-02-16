package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"brunomelo.crud/v1/model"
	"brunomelo.crud/v1/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	customerService *service.Customer
}

func (c *Customer) Post(w http.ResponseWriter, r *http.Request) {

	fmt.Println("POST")
	var model model.Customer

	json.NewDecoder(r.Body).Decode(&model)
	err := c.customerService.Save(&model)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model)
}

func (c *Customer) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	customers, err := c.customerService.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(customers)
}

func (c *Customer) FindById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GETID")
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customers, err := c.customerService.FindById(&id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if customers.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(customers)
}

func (c *Customer) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEL")
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.customerService.Delete(&id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (c *Customer) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPD")
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var model model.Customer
	err = json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	model.Id = id

	err = c.customerService.Update(&model)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model)
}

func NewCustomerRoute(service *service.Customer) *Customer {
	return &Customer{customerService: service}
}
