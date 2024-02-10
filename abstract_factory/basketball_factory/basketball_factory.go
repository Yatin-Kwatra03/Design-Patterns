package basketball_factory

type BasketballFactory struct {
	BasketballNoobPlayer        *BasketballNoobPlayer
	BasketballExperiencedPlayer *BasketballExperiencedPlayer
}

func NewBasketballFactory() *BasketballFactory {
	return &BasketballFactory{
		BasketballNoobPlayer:        &BasketballNoobPlayer{},
		BasketballExperiencedPlayer: &BasketballExperiencedPlayer{},
	}
}
