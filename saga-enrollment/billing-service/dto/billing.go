package dto

type CreateBillingRequest struct {
	OrderID   string  `json:"order_id" validate:"uuid"` // OrderID = CourseID
	StudentID string  `json:"student_id" validate:"uuid"`
	Amount    float64 `json:"amount" validate:"required,gt=0"`
}

type BillingStatus string

const (
	Pending BillingStatus = "Pending"
	Success BillingStatus = "Success"
	Failed  BillingStatus = "Failed"
)

type BillingRefusedReason string

const (
	InsufficientBalance BillingRefusedReason = "InsufficientBalance"
)

type CreateBillingResponse struct {
	ID string `json:"id"`
	CreateBillingRequest
	Status        BillingStatus        `json:"status"`
	RefusedReason BillingRefusedReason `json:"refused_reason"`
}

type GetBillingRequest struct {
	ID string `param:"id" validate:"uuid"`
}

type GetBillingResponse struct {
	ID            string               `json:"id"`
	OrderID       string               `json:"order_id"`
	StudentID     string               `json:"student_id"`
	Amount        float64              `json:"amount"`
	Status        BillingStatus        `json:"status"`
	RefusedReason BillingRefusedReason `json:"refused_reason"` // Is only set when Status is Failed
}

type GetAllBillingRequest struct {
	StudentID string `query:"student_id" validate:"omitempty,uuid"`
	OrderID   string `query:"order_id" validate:"omitempty,uuid"`
}

type GetAllBillingResponse struct {
	Billings []*GetBillingResponse `json:"billings"`
}
