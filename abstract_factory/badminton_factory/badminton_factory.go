package badminton_factory

type BadmintonFactory struct {
	BadmintonNoobPlayer        *BadmintonNoobPlayer
	BadmintonExperiencedPlayer *BadmintonExperiencedPlayer
}

func NewBadmintonFactory() *BadmintonFactory {
	return &BadmintonFactory{
		BadmintonNoobPlayer:        &BadmintonNoobPlayer{},
		BadmintonExperiencedPlayer: &BadmintonExperiencedPlayer{},
	}
}
