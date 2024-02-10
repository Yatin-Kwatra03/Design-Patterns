package abstract_factory

import (
	"github.com/personal-projects/Design-Patterns/abstract_factory/badminton_factory"
	"github.com/personal-projects/Design-Patterns/abstract_factory/basketball_factory"
)

type NoobPlayer interface {
	GetNoobPlayer() string
}

type ExperiencedPlayer interface {
	GetExperiencedPlayer() string
}

// compile time to make sure all methods of interface are implemented
var _ NoobPlayer = &basketball_factory.BasketballNoobPlayer{}
var _ ExperiencedPlayer = &basketball_factory.BasketballExperiencedPlayer{}
var _ NoobPlayer = &badminton_factory.BadmintonNoobPlayer{}
var _ ExperiencedPlayer = &badminton_factory.BadmintonExperiencedPlayer{}

// Players will act as the family of related things (noob and experienced players are related)
type Players struct {
	NoobPlayer        NoobPlayer
	ExperiencedPlayer ExperiencedPlayer
}
