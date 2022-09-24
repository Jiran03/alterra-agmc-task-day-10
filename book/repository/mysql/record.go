package repoMySQL

import (
	"github.com/Jiran03/agmc/task/day5/book/domain"
)

type Book struct {
	ID              int
	Title           string
	Writer          string
	Publisher       string
	PublicationYear int
	CreatedAt       string
	UpdatedAt       string
}

func toDomain(rec Book) domain.Book {
	return domain.Book{
		ID:              rec.ID,
		Title:           rec.Title,
		Writer:          rec.Writer,
		Publisher:       rec.Publisher,
		PublicationYear: rec.PublicationYear,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Book) Book {
	return Book{
		ID:              rec.ID,
		Title:           rec.Title,
		Writer:          rec.Writer,
		Publisher:       rec.Publisher,
		PublicationYear: rec.PublicationYear,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}
