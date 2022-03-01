package data

import (
	"database/sql"
	"log"

	"brunomelo.crud/v1/pkg/model"
)

type SqlProvider interface {
	GetConnection() *sql.DB
}

type SqlHandler struct {
	sqlProvider SqlProvider
}

func (c *SqlHandler) CustomerInsert(model *model.Customer) error {

	database := c.sqlProvider.GetConnection()
	defer database.Close()

	sql := `INSERT INTO customer ("name", birthdate, phonenumber, email, zipcode, city, street, country) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) 
			RETURNING id`

	var id int64

	if err := database.QueryRow(sql,
		model.Name,
		model.BirthDate,
		model.Contact.PhoneNumber,
		model.Contact.Email,
		model.Contact.Email,
		model.Address.ZipCode,
		model.Address.Street,
		model.Address.Country,
	).Scan(&id); err == nil {
		model.Id = id
		return nil
	} else {
		log.Fatal(err)
		return err
	}
}

func (c *SqlHandler) CustomerUpdate(model *model.Customer) error {

	database := c.sqlProvider.GetConnection()
	defer database.Close()

	sql := `UPDATE customer SET "name"=$1, birthdate=$2, phonenumber=$3, email=$4, zipcode=$5, city=$6, street=$7, country=$8 WHERE id=$9`

	if _, err := database.Exec(sql,
		model.Name,
		model.BirthDate,
		model.Contact.PhoneNumber,
		model.Contact.Email,
		model.Contact.Email,
		model.Address.ZipCode,
		model.Address.Street,
		model.Address.Country,
		model.Id,
	); err == nil {
		return nil
	} else {
		return err
	}
}

func (c *SqlHandler) GetAllWithoutPagination() ([]model.Customer, error) {

	database := c.sqlProvider.GetConnection()
	defer database.Close()

	sql := `SELECT id, "name", birthdate, phonenumber, email, zipcode, city, street, country FROM customer`

	rows, err := database.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var models []model.Customer

	for rows.Next() {
		var model model.Customer

		err = rows.Scan(&model.Id, &model.Name, &model.BirthDate, &model.Contact.PhoneNumber, &model.Contact.Email, &model.Address.ZipCode, &model.Address.City, &model.Address.Street, &model.Address.Country)

		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}

func (c *SqlHandler) CustomerDelete(id *int64) error {

	database := c.sqlProvider.GetConnection()
	defer database.Close()

	sql := `DELETE FROM customer WHERE id = $1`

	res, err := database.Exec(sql, id)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func (c *SqlHandler) CustomerFindById(id *int64) (model.Customer, error) {

	database := c.sqlProvider.GetConnection()
	defer database.Close()

	sql := `SELECT id, "name", birthdate, phonenumber, email, zipcode, city, street, country FROM customer WHERE id = $1`

	rows, err := database.Query(sql, id)

	var model model.Customer

	if err != nil {
		return model, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&model.Id, &model.Name, &model.BirthDate, &model.Contact.PhoneNumber, &model.Contact.Email, &model.Address.ZipCode, &model.Address.City, &model.Address.Street, &model.Address.Country)

		if err != nil {
			return model, err
		}
	}

	return model, nil
}

func NewSqlHanlder(provider SqlProvider) *SqlHandler {
	return &SqlHandler{sqlProvider: provider}
}
