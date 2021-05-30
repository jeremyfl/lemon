package model

// Customer is the main customer model
type Customer struct {
	ID        int64  `db:"pk" column:"tbl_id" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
}
