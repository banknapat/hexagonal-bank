package repository
// ดึง database และ จัดการ database
type Customer struct {
	CustonerID int `db:"customer_id"`
	Name string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City string `db:"city"`
	ZipCode string `db:"zip_code"`
	Status string `db:"status"`
}


type Repository interface {
	GetAll()
	GetById(int)
}