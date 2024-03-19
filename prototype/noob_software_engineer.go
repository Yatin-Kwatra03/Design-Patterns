package prototype

// noobSoftwareEngineer: concrete implementation
// of the software engineer
type noobSoftwareEngineer struct {
	id         int
	name       string
	experience int
	skill      string
}

// compile time safety check
var _ softwareEngineer = &noobSoftwareEngineer{}

//   - we can argue on id can't be same for two software engineers
//     so, we can add random id creation logic inside newNoobSoftwareEngineer
//
// - but then we can argue its not a proper clone if we values are different.
//
//	So keeping id as a parameter to keep things simple.
func newNoobSoftwareEngineer(id int, name string, experience int) *noobSoftwareEngineer {
	return &noobSoftwareEngineer{
		id:         id,
		name:       name,
		experience: experience,
		skill:      "low",
	}
}

func (s *noobSoftwareEngineer) GetId() int {
	return s.id
}

func (s *noobSoftwareEngineer) Clone() softwareEngineer {
	return newNoobSoftwareEngineer(s.id, s.name, s.experience)
}
