package repoMySQL

import (
	"github.com/Jiran03/agmc/task/day5/book/domain"
)

type bookRepository struct {
	DB Book
}

var books = []Book{
	{
		ID:              1,
		Title:           "title 1",
		Writer:          "writer 1",
		Publisher:       "publisher 1",
		PublicationYear: 2022,
		CreatedAt:       "",
		UpdatedAt:       "",
	},
	{
		ID:              2,
		Title:           "title 2",
		Writer:          "writer 2",
		Publisher:       "publisher 2",
		PublicationYear: 2022,
		CreatedAt:       "",
		UpdatedAt:       "",
	},
}

// Update implements domain.Repository
func (ur bookRepository) Update(domain domain.Book) (bookObj domain.Book, err error) {
	var newRecord Book
	// newBook := new(Book)
	rec := fromDomain(domain)
	for i, v := range books {
		if rec.ID == v.ID {
			newRecord = books[i]
		}
	}

	newRecord.Title = rec.Title
	newRecord.Writer = rec.Writer
	newRecord.Publisher = rec.Publisher
	newRecord.PublicationYear = rec.PublicationYear

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (ur bookRepository) GetByID(id int) (domain domain.Book, err error) {
	var record Book
	for i, v := range books {
		if id == v.ID {
			record = books[i]
		}
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (ur bookRepository) Get() (bookObj []domain.Book, err error) {
	for _, value := range books {
		bookObj = append(bookObj, toDomain(value))
	}

	return bookObj, nil
}

// Create implements domain.Repository
func (ur bookRepository) Create(domain domain.Book) (bookObj domain.Book, err error) {
	domain.ID = len(books) + 1
	record := fromDomain(domain)
	bookObj = toDomain(record)

	return bookObj, nil
}

// Delete implements domain.Repository
func (ur bookRepository) Delete(id int) (err error) {
	arrRecord := make([]Book, 0)
	for i, v := range books {
		if id == v.ID {
			arrRecord = append(arrRecord, books[:i]...)
			arrRecord = append(arrRecord, books[i+1:]...)
			books = arrRecord
		}
	}

	return nil
}

func NewBookRepository(db Book) domain.Repository {
	return bookRepository{
		DB: db,
	}
}
