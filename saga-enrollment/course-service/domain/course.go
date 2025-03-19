package domain

type Course struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Seats    int      `json:"seats"`
	Students []string `json:"students"`
}
