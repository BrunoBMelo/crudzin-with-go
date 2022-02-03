package ioc

import (
	"brunomelo.crud/v1/config"
	"brunomelo.crud/v1/data"
	"brunomelo.crud/v1/service"
)

type containerDI struct {
	CustomerServices *service.Customer
	Config           *config.Config
}

func (di *containerDI) register() {

	di.Config = config.Load()
	sqlProvider := data.NewPostgreSql(di.Config.ProviderName, di.Config.StringConnection)
	sqlHandler := data.NewSqlHanlder(sqlProvider)
	di.CustomerServices = service.NewCustomerService(sqlHandler)
}

func NewContainerDI() *containerDI {
	di := &containerDI{}
	di.register()
	return di
}
