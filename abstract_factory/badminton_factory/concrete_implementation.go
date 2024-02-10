package badminton_factory

type BadmintonNoobPlayer struct{}
type BadmintonExperiencedPlayer struct{}

func (s *BadmintonNoobPlayer) GetNoobPlayer() string {
	return "created a noob player of badminton"
}

func (s *BadmintonExperiencedPlayer) GetExperiencedPlayer() string {
	return "created a experienced player of badminton"
}
