package domain

type CourseService interface {
	Create(dto CreateCourseRequest) (*Course, error)
	FindByID(id string) (*Course, error)
}

type CourseRepository interface {
	Create(dto CreateCourseRequest) (*Course, error)
	FindByID(id string) (*Course, error)
}
