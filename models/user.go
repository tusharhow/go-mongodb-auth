package models


type User struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}
