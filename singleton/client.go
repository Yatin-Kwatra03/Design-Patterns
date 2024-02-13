package singleton

func GetCricketTeamInstance() *CricketTeam {
	return newCricketTeam()
}
