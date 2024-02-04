package factory

type Biopic struct {
	id    string
	name  string
	refNo string
}

var _ Book = &Biopic{}

func NewBiopic(name string) *Biopic {
	return &Biopic{
		id:    "1234",
		name:  name,
		refNo: "687",
	}
}

func (s *Biopic) GetBookName() string {
	return s.name
}

func (s *Biopic) GetBookId() string {
	return s.id
}
