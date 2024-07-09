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
	rows, err := c.db.Query("SELECT * FROM cards WHERE type = 2 LIMIT 10")
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
