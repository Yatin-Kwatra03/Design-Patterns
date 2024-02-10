package basketball_factory

type BasketballNoobPlayer struct{}
type BasketballExperiencedPlayer struct{}

func (s *BasketballNoobPlayer) GetNoobPlayer() string {
	return "created a noob player of Basketball"
}

func (s *BasketballExperiencedPlayer) GetExperiencedPlayer() string {
	return "created a experienced player of Basketball"
}
