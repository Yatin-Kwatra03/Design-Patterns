package prototype

// experiencedSoftwareEngineer: concrete implementation
// of the software engineer
type experiencedSoftwareEngineer struct {
	id         int
	name       string
	experience int
	skill      string
}

// compile time safety check
var _ softwareEngineer = &experiencedSoftwareEngineer{}

//   - we can argue on id can't be same for two software engineers
//     so, we can add random id creation logic inside newExperiencedSoftwareEngineer
//
// - but then we can argue its not a proper clone if we values are different.
//
//	So keeping id as a parameter to keep things simple.
func newExperiencedSoftwareEngineer(id int, name string, experience int) *experiencedSoftwareEngineer {
	return &experiencedSoftwareEngineer{
		id:         id,
		name:       name,
		experience: experience,
		skill:      "high",
	}
}

func (s *experiencedSoftwareEngineer) GetId() int {
	return s.id
}

func (s *experiencedSoftwareEngineer) Clone() softwareEngineer {
	return newExperiencedSoftwareEngineer(s.id, s.name, s.experience)
}
