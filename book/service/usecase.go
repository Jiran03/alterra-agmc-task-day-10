package service

import (
	"github.com/Jiran03/agmc/task/day5/book/domain"
	timeHelper "github.com/Jiran03/agmc/task/day5/helpers/time"
)

type bookService struct {
	repository domain.Repository
}

// UpdateData implements domain.Service
func (bs bookService) UpdateData(id int, domain domain.Book) (bookObj domain.Book, err error) {
	if bookObj, err = bs.GetByID(id); err != nil {
		return bookObj, err
	}

	domain.ID = id
	if bookObj, err = bs.repository.Update(domain); err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// GetByID implements domain.Service
func (bs bookService) GetByID(id int) (bookObj domain.Book, err error) {
	bookObj, err = bs.repository.GetByID(id)
	if err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

func (bs bookService) InsertData(domain domain.Book) (bookObj domain.Book, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	bookObj, err = bs.repository.Create(domain)
	if err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// GetAllData implements domain.Service
func (bs bookService) GetAllData() (bookObj []domain.Book, err error) {
	bookObj, _ = bs.repository.Get()
	if err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// DeleteData implements domain.Service
func (bs bookService) DeleteData(id int) (err error) {
	errResp := bs.repository.Delete(id)
	if errResp != nil {
		return errResp
	}

	return nil
}

func NewBookService(repo domain.Repository) domain.Service {
	return bookService{
		repository: repo,
	}
}
