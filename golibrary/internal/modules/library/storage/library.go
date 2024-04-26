package storage

import (
	"context"
	"golibrary/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LibraryRepository interface {
	CreateBook(ctx context.Context, book models.Book) (int, error)
	CreateAuthor(ctx context.Context, author models.Author) (int, error)
	CreateAuthorBook(ctx context.Context, id_author, id_book int) error

	GetBookById(ctx context.Context, id_book int) (models.Book, error)
	GetAuthorById(ctx context.Context, id_author int) (models.Author, error)

	ListBook(ctx context.Context) ([]models.BookWithAuthor, error)
	ListAuthor(ctx context.Context) ([]models.AuthorWithBooks, error)

	CreateUserBook(ctx context.Context, id_user, id_book int) error
	DeleteUserBook(ctx context.Context, id_user, id_book int) error

	GetUserIdByUser(ctx context.Context, id_user int) (models.User, error)
	RentedBooks(ctx context.Context) ([]models.UsersWithBook, error)
}
type LibraryStorage struct {
	pool *pgxpool.Pool
}

func NewLibraryStorage(pool *pgxpool.Pool) LibraryRepository {
	return &LibraryStorage{pool}
}

func (l *LibraryStorage) CreateBook(ctx context.Context, book models.Book) (int, error) {
	var id int

	sql := "INSERT INTO book (name_book) VALUES ($1) RETURNING id"
	err := l.pool.QueryRow(ctx, sql, book.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (l *LibraryStorage) CreateAuthor(ctx context.Context, author models.Author) (int, error) {
	var id int

	sql := "INSERT INTO author (name_author) VALUES ($1) RETURNING id"
	err := l.pool.QueryRow(ctx, sql, author.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (l *LibraryStorage) CreateAuthorBook(ctx context.Context, id_author, id_book int) error {
	sql := "INSERT INTO author_book (id_author, id_book) VALUES ($1, $2)"
	row, err := l.pool.Query(ctx, sql, id_author, id_book)
	if err != nil {
		return err
	}
	defer row.Close()

	return err
}

func (l *LibraryStorage) GetBookById(ctx context.Context, id_book int) (models.Book, error) {
	var book models.Book

	sql := "SELECT name_book FROM book WHERE id=$1"
	err := l.pool.QueryRow(ctx, sql, id_book).Scan(&book.Name)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (l *LibraryStorage) GetAuthorById(ctx context.Context, id_author int) (models.Author, error) {
	var author models.Author

	sql := "SELECT name_author FROM author WHERE id=$1"
	err := l.pool.QueryRow(ctx, sql, id_author).Scan(&author.Name)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (l *LibraryStorage) ListBook(ctx context.Context) ([]models.BookWithAuthor, error) {
	books := make([]models.BookWithAuthor, 0, 1)

	sql := "SELECT id_author, id_book FROM author_book"
	rows, err := l.pool.Query(ctx, sql)
	if err != nil {
		return []models.BookWithAuthor{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_author, id_book int

		if err = rows.Scan(&id_author, &id_book); err != nil {
			return books, err
		}
		author, err := l.GetAuthorById(ctx, id_author)
		if err != nil {
			return books, err
		}
		book, err := l.GetBookById(ctx, id_book)
		if err != nil {
			return books, err
		}
		books = append(books, models.BookWithAuthor{
			Author: author,
			Book:   book,
		})
	}
	return books, nil
}

func (l *LibraryStorage) ListAuthor(ctx context.Context) ([]models.AuthorWithBooks, error) {
	authors := make([]models.AuthorWithBooks, 0)

	sql := "SELECT id_author, id_book FROM author_book ORDER BY id_author ASC"
	rows, err := l.pool.Query(ctx, sql)
	if err != nil {
		return []models.AuthorWithBooks{}, err
	}
	defer rows.Close()

	i := -1
	prev_id_author := 0
	for rows.Next() {
		var id_author, id_book int

		if err = rows.Scan(&id_author, &id_book); err != nil {
			return authors, err
		}
		author, err := l.GetAuthorById(ctx, id_author)
		if err != nil {
			return authors, err
		}
		book, err := l.GetBookById(ctx, id_book)
		if err != nil {
			return authors, err
		}
		if prev_id_author != id_author {
			prev_id_author = id_author
			books := make([]models.Book, 0)
			books = append(books, book)
			authors = append(authors, models.AuthorWithBooks{
				Author: author,
				Books:  books,
			})
			i++
		} else {
			authors[i].Books = append(authors[i].Books, book)
		}
	}
	return authors, nil
}

func (l *LibraryStorage) CreateUserBook(ctx context.Context, id_user, id_book int) error {
	sql := "INSERT INTO user_book (id_user, id_book) VALUES ($1, $2)"
	row, err := l.pool.Query(ctx, sql, id_user, id_book)
	if err != nil {
		return err
	}
	defer row.Close()

	return err
}

func (l *LibraryStorage) DeleteUserBook(ctx context.Context, id_user, id_book int) error {
	sql := "DELETE FROM user_book WHERE id_user = $1 AND id_book = $2"
	row, err := l.pool.Query(ctx, sql, id_user, id_book)
	if err != nil {
		return err
	}
	defer row.Close()

	return err
}

func (l *LibraryStorage) GetUserIdByUser(ctx context.Context, id_user int) (models.User, error) {
	var user models.User

	sql := "SELECT username FROM users WHERE id=$1"
	err := l.pool.QueryRow(ctx, sql, id_user).Scan(&user.Name)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (l *LibraryStorage) RentedBooks(ctx context.Context) ([]models.UsersWithBook, error) {
	usersWithBook := make([]models.UsersWithBook, 0)

	sql := "SELECT id_user, id_book FROM user_book ORDER BY id_user ASC"
	rows, err := l.pool.Query(ctx, sql)
	if err != nil {
		return []models.UsersWithBook{}, nil
	}
	defer rows.Close()

	i := -1
	prev_id_user := 0
	for rows.Next() {
		var id_user, id_book, id_author int

		if err = rows.Scan(&id_user, &id_book); err != nil {
			return usersWithBook, err
		}
		user, err := l.GetUserIdByUser(ctx, id_user)
		if err != nil {
			return usersWithBook, err
		}

		book, err := l.GetBookById(ctx, id_book)
		if err != nil {
			return usersWithBook, err
		}

		sql = "SELECT id_author FROM author_book WHERE id_book = $1"
		err = l.pool.QueryRow(ctx, sql, id_book).Scan(&id_author)
		if err != nil {
			return usersWithBook, err
		}
		author, err := l.GetAuthorById(ctx, id_author)
		if err != nil {
			return usersWithBook, err
		}
		if prev_id_user != id_user {
			prev_id_user = id_user
			bookWithAuthor := make([]models.BookWithAuthor, 0)
			bookWithAuthor = append(bookWithAuthor, models.BookWithAuthor{
				Author: author,
				Book:   book,
			})
			usersWithBook = append(usersWithBook, models.UsersWithBook{
				User:        user,
				RentedBooks: bookWithAuthor,
			})
			i++
		} else {
			usersWithBook[i].RentedBooks = append(usersWithBook[i].RentedBooks, models.BookWithAuthor{
				Author: author,
				Book:   book,
			})
		}
	}
	return usersWithBook, nil
}
