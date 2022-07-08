package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(Id int, updateBookRequest UpdateBookRequest) (Book, error)
	Delete(Id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, err := bookRequest.Price.Int64()
	rating, err := bookRequest.Rating.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
func (s *service) Update(Id int, updateBookRequest UpdateBookRequest) (Book, error) {
	book, err := s.repository.FindByID(Id)

	price, err := updateBookRequest.Price.Int64()
	rating, err := updateBookRequest.Rating.Int64()

	book.Title = updateBookRequest.Title
	book.Description = updateBookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}
func (s *service) Delete(Id int) (Book, error) {
	book, err := s.repository.FindByID(Id)
	deleteBook, err := s.repository.Delete(book)
	return deleteBook, err
}
