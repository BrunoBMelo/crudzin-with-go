package service

import (
	"brunomelo.crud/v1/data"
	"brunomelo.crud/v1/model"
)

type Customer struct {
	sql *data.SqlHandler
}

func (s *Customer) Save(customer *model.Customer) error {

	err := s.sql.CustomerInsert(customer)

	if err != nil {
		return err
	}

	return nil
}

func (s *Customer) GetAll() ([]model.Customer, error) {

	models, err := s.sql.GetAllWithoutPagination()

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *Customer) Delete(id *int64) error {
	return s.sql.CustomerDelete(id)
}

func (s *Customer) FindById(id *int64) (model.Customer, error) {
	return s.sql.CustomerFindById(id)
}

func (s *Customer) Update(model *model.Customer) error {
	return s.sql.CustomerUpdate(model)
}

func NewCustomerService(sql *data.SqlHandler) *Customer {
	return &Customer{sql: sql}
}
