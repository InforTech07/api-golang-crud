package models

type Person struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
