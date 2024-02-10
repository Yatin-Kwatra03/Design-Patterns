package abstract_factory

import (
	"github.com/personal-projects/Design-Patterns/abstract_factory/badminton_factory"
	"github.com/personal-projects/Design-Patterns/abstract_factory/basketball_factory"
)

type AbstractFactory struct {
	basketballFactory *basketball_factory.BasketballFactory
	badmintonFactory  *badminton_factory.BadmintonFactory
}

func NewAbstractFactory() *AbstractFactory {
	basketballFactory := basketball_factory.NewBasketballFactory()
	badmintonFactory := badminton_factory.NewBadmintonFactory()

	return &AbstractFactory{
		basketballFactory: basketballFactory,
		badmintonFactory:  badmintonFactory,
	}
}

func ParticipateAsPerTheSportsSkill(sportsSkill string) *Players {
	abstractFactoryService := initializeAbstractFactoryService()

	if sportsSkill == "basketball" {
		return &Players{
			NoobPlayer:        abstractFactoryService.basketballFactory.BasketballNoobPlayer,
			ExperiencedPlayer: abstractFactoryService.basketballFactory.BasketballExperiencedPlayer,
		}
	}

	return &Players{
		NoobPlayer:        abstractFactoryService.badmintonFactory.BadmintonNoobPlayer,
		ExperiencedPlayer: abstractFactoryService.badmintonFactory.BadmintonExperiencedPlayer,
	}
}

func initializeAbstractFactoryService() *AbstractFactory {
	basketballFactory := basketball_factory.NewBasketballFactory()
	badmintonFactory := badminton_factory.NewBadmintonFactory()

	return &AbstractFactory{
		basketballFactory: basketballFactory,
		badmintonFactory:  badmintonFactory,
	}
}
