package types

type Line struct {
	Id    int
	Type  int
	Front string
	Back  string
	Known bool
}

type Lines = []Line

type Data struct {
	Lines Lines
}

type CardsInterface interface {
	GetCards() (Data, error)
	GetCardsFilter(string) ([]Line, error)
	GetCardsById(string) (Line, error)
	GetCardAlredyKnow(string) (Line, error)
	MarkUnknown(string) error
}
