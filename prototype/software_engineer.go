package prototype

// softwareEngineer : interface to be implemented
// by each class which wants to support cloning.
type softwareEngineer interface {
	GetId() int
	Clone() softwareEngineer
}
