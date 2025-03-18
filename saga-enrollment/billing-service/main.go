package main

import (
	"github.com/do4-mc-homework/labs/saga-enrollment/billing-service/api"
	repository "github.com/do4-mc-homework/labs/saga-enrollment/billing-service/repositories"
	"github.com/do4-mc-homework/labs/saga-enrollment/billing-service/service"
)

func main() {
	repository := repository.NewInMemoryBillingRepository()
	service := service.NewBillingService(repository)
	server := api.NewBillingApi(service)
	server.Start(":8081")
}
