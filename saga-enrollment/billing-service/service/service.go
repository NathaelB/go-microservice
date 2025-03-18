package service

import (
	"math/rand"

	"github.com/do4-mc-homework/labs/saga-enrollment/billing-service/dto"
	repository "github.com/do4-mc-homework/labs/saga-enrollment/billing-service/repositories"
	"github.com/google/uuid"
)

type BillingService interface {
	Save(request *dto.CreateBillingRequest) (*dto.CreateBillingResponse, error)
	Get(request *dto.GetBillingRequest) (*dto.GetBillingResponse, error)
	GetAll(request *dto.GetAllBillingRequest) ([]*dto.GetBillingResponse, error)
}

type ApiBillingService struct {
	repository repository.BillingRepository
}

func NewBillingService(repository repository.BillingRepository) *ApiBillingService {
	return &ApiBillingService{
		repository: repository,
	}
}

func (api *ApiBillingService) GetAll(request *dto.GetAllBillingRequest) ([]*dto.GetBillingResponse, error) {
	filter := repository.Filter{
		StudentID: request.StudentID,
		OrderID:   request.OrderID,
	}
	billings, err := api.repository.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var response []*dto.GetBillingResponse
	for _, billing := range billings {
		response = append(response, &dto.GetBillingResponse{
			ID:            billing.ID,
			OrderID:       billing.OrderID,
			Amount:        billing.Amount,
			StudentID:     billing.StudentID,
			Status:        dto.BillingStatus(billing.Status),
			RefusedReason: dto.BillingRefusedReason(billing.RefusedReason),
		})
	}

	return response, nil
}

func (api *ApiBillingService) Get(request *dto.GetBillingRequest) (*dto.GetBillingResponse, error) {
	billing, err := api.repository.Get(request.ID)
	if err != nil {
		return nil, err
	}

	if billing == nil {
		return nil, nil
	}

	return &dto.GetBillingResponse{
		ID:            billing.ID,
		OrderID:       billing.OrderID,
		Amount:        billing.Amount,
		StudentID:     billing.StudentID,
		Status:        dto.BillingStatus(billing.Status),
		RefusedReason: dto.BillingRefusedReason(billing.RefusedReason),
	}, nil
}

func (api *ApiBillingService) Save(request *dto.CreateBillingRequest) (*dto.CreateBillingResponse, error) {

	status := string(dto.Pending)
	refusedReason := ""

	random := rand.Float64()
	// For testing purposes, we randomly set the status
	if random < 0.33 {
		status = string(dto.Failed)
		refusedReason = string(dto.InsufficientBalance)
	}
	if random > 0.66 {
		status = string(dto.Success)
	}

	id := uuid.New().String()
	err := api.repository.Save(&repository.Billing{
		ID:            id,
		OrderID:       request.OrderID,
		StudentID:     request.StudentID,
		Amount:        request.Amount,
		Status:        status,
		RefusedReason: refusedReason,
	})
	if err != nil {
		return nil, err
	}
	return &dto.CreateBillingResponse{
		ID:                   id,
		CreateBillingRequest: *request,
		Status:               dto.BillingStatus(status),
		RefusedReason:        dto.BillingRefusedReason(refusedReason),
	}, nil
}
