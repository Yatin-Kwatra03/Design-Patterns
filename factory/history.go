package factory

type History struct{}

func NewHistory() *History {
	return &History{}
}

func (s *History) GetBook() (*Book, error) {
	return &Book{
		Id:       "23",
		Name:     "yatin dae raaz",
		BookType: BookType_BOOK_TYPE_HISTORY,
		Price:    12345,
	}, nil
}
