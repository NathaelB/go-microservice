package repository

type Billing struct {
	ID            string
	OrderID       string
	StudentID     string
	Amount        float64
	Status        string
	RefusedReason string
}

type Filter struct {
	StudentID string
	OrderID   string
}

type BillingRepository interface {
	Save(billing *Billing) error
	Get(id string) (*Billing, error)
	GetAll(filter Filter) ([]*Billing, error)
}

type InMemoryBillingRepository struct {
	billings map[string]*Billing
}

func NewInMemoryBillingRepository() *InMemoryBillingRepository {
	return &InMemoryBillingRepository{
		billings: make(map[string]*Billing),
	}
}

// This function is used to get billings based on the filter
func (r *InMemoryBillingRepository) GetAll(filter Filter) ([]*Billing, error) {
	var billings []*Billing
	for _, billing := range r.billings {

		if filter.StudentID == "" && filter.OrderID == "" {
			billings = append(billings, billing)
			continue
		}

		if filter.StudentID != "" && billing.StudentID != filter.StudentID {
			continue
		}

		if filter.OrderID != "" && billing.OrderID != filter.OrderID {
			continue
		}

		billings = append(billings, billing)
	}

	return billings, nil
}

func (r *InMemoryBillingRepository) Save(billing *Billing) error {
	r.billings[billing.ID] = billing
	return nil
}

func (r *InMemoryBillingRepository) Get(id string) (*Billing, error) {
	billing, ok := r.billings[id]
	if !ok {
		return nil, nil
	}

	return billing, nil
}
