package factory

type Biopic struct{}

func NewBiopic() *Biopic {
	return &Biopic{}
}

func (s *Biopic) GetBook() (*Book, error) {
	return &Book{
		Id:       "23",
		Name:     "yatin dae basketball dae raaz",
		BookType: BookType_BOOK_TYPE_BIOPIC,
		Price:    12345,
	}, nil
}
