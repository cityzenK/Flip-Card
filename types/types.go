package types

type Line struct {
	Id    int
	Type  int
	Front string
	Back  string
	Known bool
}

type CardsInterface interface {
	GetCards() ([]Line, error)
	GetCardsFilter(string) ([]Line, error)
	GetCardsById(string) (Line, error)
	GetCardAlredyKnow(string) (Line, error)
	MarkUnknown(string) error
}
