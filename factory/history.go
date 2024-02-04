package factory

type History struct {
	id   string
	name string
}

var _ Book = &Biopic{}

func NewHistory(name string) *History {
	return &History{
		id:   "345",
		name: name,
	}
}

func (s *History) GetBookName() string {
	return s.name
}

func (s *History) GetBookId() string {
	return s.id
}
