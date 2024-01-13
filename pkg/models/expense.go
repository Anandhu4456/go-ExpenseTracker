package models

type Section struct {
	// ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserName    string  `json:"user_name"`
	Buget       float64 `json:"buget"`
	Description string  `json:"description"`
}
