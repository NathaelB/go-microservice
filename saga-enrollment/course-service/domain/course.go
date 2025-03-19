package domain

type Course struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Seats    int       `json:"seats"`
	Students []Student `json:"students" gorm:"many2many:course_students;"`
}

type Student struct {
	ID string `json:"id"`
}

// CourseStudent represents a many-to-many relationship between courses and students
type CourseStudent struct {
	CourseID  string `json:"course_id" gorm:"primaryKey"`
	StudentID string `json:"student_id" gorm:"primaryKey"`
}
