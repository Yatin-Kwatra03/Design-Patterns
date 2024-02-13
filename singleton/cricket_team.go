package singleton

type CricketTeam struct {
	Name string
}

var (
	cricketTeamInstance *CricketTeam
	InstanceNo          int
)

// getInstance private method to create instance of a class
// This doesn't really restrict the outside services to instantiate the cricket team instance because in go
// we don't have a constructor concept, anyone can just instantiate &CricketTeam{}. But it we follow the logic
// to always go to the domain service to get the instance, then it can be implemented in this way.
func newCricketTeam() *CricketTeam {
	if cricketTeamInstance == nil {
		cricketTeamInstance = &CricketTeam{
			Name: "I am instantiated",
		}
		InstanceNo++
	}
	return cricketTeamInstance
}
