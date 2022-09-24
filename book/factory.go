package book

import (
	handlerAPI "github.com/Jiran03/agmc/task/day5/book/handler/api"
	repoMySQL "github.com/Jiran03/agmc/task/day5/book/repository/mysql"
	service "github.com/Jiran03/agmc/task/day5/book/service"
)

func NewBookFactory(db repoMySQL.Book) (bookHandler handlerAPI.BookHandler) {
	bookRepo := repoMySQL.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler = handlerAPI.NewBookHandler(bookService)
	return
}
