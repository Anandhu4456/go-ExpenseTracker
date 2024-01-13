package models

type Section struct {
	// ID          primitive.ObjectID
	UserName    string  `json:"user_name"`
	Buget       float64 `json:"buget"`
	Description string  `json:"description"`
}
