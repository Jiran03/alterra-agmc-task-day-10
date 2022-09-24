package domain

type Service interface {
	InsertData(domain Book) (bookObj Book, err error)
	UpdateData(id int, domain Book) (bookObj Book, err error)
	GetAllData() (bookObj []Book, err error)
	GetByID(id int) (bookObj Book, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Book) (bookObj Book, err error)
	Update(domain Book) (bookObj Book, err error)
	Get() (bookObj []Book, err error)
	GetByID(id int) (domain Book, err error)
	Delete(id int) (err error)
}
