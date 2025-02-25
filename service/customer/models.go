package customer

type Customer struct {
	ID   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
	Age  int    `gorm:"age" json:"age"`
}

type Customers []Customer
