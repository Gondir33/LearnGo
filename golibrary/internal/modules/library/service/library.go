package service

import (
	"context"
	"golibrary/internal/models"
	"golibrary/internal/modules/library/storage"
)

type LibraryServicere interface {
	CreateBook(ctx context.Context, book models.Book, id_author int) error
	RentedBooks(ctx context.Context) ([]models.UsersWithBook, error)
	CreateAuthor(ctx context.Context, author models.Author) error
	ListBook(ctx context.Context) ([]models.BookWithAuthor, error)
	ListAuthor(ctx context.Context) ([]models.AuthorWithBooks, error)
	GetBook(ctx context.Context, id_user, id_book int) error
	HandBook(ctx context.Context, id_user, id_book int) error
}

type LibraryService struct {
	storage.LibraryRepository
}

func NewUserService(LibraryRepository storage.LibraryRepository) LibraryServicere {
	return &LibraryService{
		LibraryRepository: LibraryRepository,
	}
}

func (l *LibraryService) CreateAuthor(ctx context.Context, author models.Author) error {
	_, err := l.LibraryRepository.CreateAuthor(ctx, author)
	return err
}

func (l *LibraryService) CreateBook(ctx context.Context, book models.Book, id_author int) error {
	id_book, err := l.LibraryRepository.CreateBook(ctx, book)
	if err != nil {
		return nil
	}
	err = l.CreateAuthorBook(ctx, id_author, id_book)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) ListBook(ctx context.Context) ([]models.BookWithAuthor, error) {
	return l.LibraryRepository.ListBook(ctx)
}

func (l *LibraryService) ListAuthor(ctx context.Context) ([]models.AuthorWithBooks, error) {
	return l.LibraryRepository.ListAuthor(ctx)
}
func (l *LibraryService) GetBook(ctx context.Context, id_user, id_book int) error {
	return l.LibraryRepository.CreateUserBook(ctx, id_user, id_book)
}
func (l *LibraryService) HandBook(ctx context.Context, id_user, id_book int) error {
	return l.LibraryRepository.DeleteUserBook(ctx, id_user, id_book)
}
func (l *LibraryService) RentedBooks(ctx context.Context) ([]models.UsersWithBook, error) {
	return l.LibraryRepository.RentedBooks(ctx)
}
