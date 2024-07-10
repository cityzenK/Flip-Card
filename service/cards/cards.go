package cards

import (
	"cards/types"
	"database/sql"
)

type Cards struct {
	db *sql.DB
}

func NewCards(db *sql.DB) *Cards {
	return &Cards{db: db}
}

func (c *Cards) GetCards() ([]types.Line, error) {
	cards := []types.Line{}
	rows, err := c.db.Query("SELECT * FROM cards")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var line types.Line
		if err := rows.Scan(&line.Id, &line.Type, &line.Front, &line.Back, &line.Known); err != nil {
			return nil, err
		}
		cards = append(cards, line)
	}

	rows.Close()
	return cards, nil
}

func (c *Cards) GetCardsFilter(filter string) ([]types.Line, error) {
	query := "SELECT * FROM cards " + filter
	cards := []types.Line{}
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var line types.Line
		if err := rows.Scan(&line.Id, &line.Type, &line.Front, &line.Back, &line.Known); err != nil {
			return nil, err
		}
		cards = append(cards, line)
	}

	return cards, nil
}
