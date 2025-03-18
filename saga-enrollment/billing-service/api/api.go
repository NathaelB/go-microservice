package api

import (
	"net/http"

	"github.com/do4-mc-homework/labs/saga-enrollment/billing-service/dto"
	"github.com/do4-mc-homework/labs/saga-enrollment/billing-service/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BillingApi struct {
	server  *echo.Echo
	service service.BillingService
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewBillingApi(service service.BillingService) *BillingApi {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	return &BillingApi{
		service: service,
		server:  e,
	}
}

func (api *BillingApi) save(c echo.Context) error {
	createBilling := new(dto.CreateBillingRequest)
	if err := c.Bind(createBilling); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(createBilling); err != nil {
		return err
	}
	billing, err := api.service.Save(createBilling)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, billing)
}

func (api *BillingApi) get(c echo.Context) error {
	getBilling := new(dto.GetBillingRequest)
	if err := c.Bind(getBilling); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(getBilling); err != nil {
		return err
	}
	billing, err := api.service.Get(getBilling)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if billing == nil {
		return c.String(http.StatusNotFound, "Billing not found")
	}
	return c.JSON(http.StatusOK, billing)
}

func (api *BillingApi) getAll(c echo.Context) error {
	filter := new(dto.GetAllBillingRequest)
	if err := c.Bind(filter); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(filter); err != nil {
		return err
	}
	billings, err := api.service.GetAll(filter)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if len(billings) == 0 {
		return c.String(http.StatusNotFound, "Billing not found")
	}
	return c.JSON(http.StatusOK, billings)
}

// Define the API routes here
func (api *BillingApi) RegisterRoutes() {
	api.server.POST("/billing", api.save)
	api.server.GET("/billing/:id", api.get)
	api.server.GET("/billing", api.getAll)
}

func (api *BillingApi) Start(port string) {
	api.RegisterRoutes()
	api.server.Logger.Fatal(api.server.Start(port))
}
