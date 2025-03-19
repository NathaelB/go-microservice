package domain

type CreateCourseRequest struct {
	Title string `json:"title"`
	Seats int    `json:"seats"`
}


