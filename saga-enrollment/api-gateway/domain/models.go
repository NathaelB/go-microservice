package domain

type Enrollment struct {
	ID            string `json:"id" gorm:"primaryKey;default:concat('ENR', floor(random() * 9000 + 1000)::text)"`
	StudentID     string `json:"student_id"`
	CourseID      string `json:"course_id"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason"`
	PaymentMethod string `json:"payment_method"`
}

type EnrollmentStatus string

const (
	Pending EnrollmentStatus = "Pending"
	Success EnrollmentStatus = "Success"
	Failed  EnrollmentStatus = "Failed"
)

type EnrollmentFailureReason string

const (
	InsufficientBalance EnrollmentFailureReason = "InsufficientBalance"
)
