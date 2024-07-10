package cards

import (
	"cards/types"
	"database/sql"
	"log"
)

type Cards struct {
	db *sql.DB
}

func NewCards(db *sql.DB) *Cards {
	return &Cards{db: db}
}

func (c *Cards) GetCards() ([]types.Line, error) {
	cards := []types.Line{}
	rows, err := c.db.Query("SELECT * FROM cards ORDER BY id DESC")
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

func (c *Cards) GetCardsById(cardId string) (types.Line, error) {
	var card types.Line

	query := "SELECT id, type, front, back, known FROM cards where id = " + cardId + " LIMIT 1"

	row, err := c.db.Query(query)
	if err != nil {
		return types.Line{}, err
	}
	if err := row.Scan(&card.Id, &card.Type, &card.Front, &card.Back, &card.Known); err != nil {
		return types.Line{}, err
	}

	return card, nil
}

func (c *Cards) GetCardAlredyKnow(card_type string) (types.Line, error) {
	var card types.Line

	query := "SELECT id, type, front, back ,known FROM  cards WHERER type =" + card_type + " AND know = 1 ORDER BY RANDOM() LIMIT 1"

	row, err := c.db.Query(query)
	if err != nil {
		return types.Line{}, err
	}
	if err := row.Scan(&card.Id, &card.Type, &card.Front, &card.Back, &card.Known); err != nil {
		return types.Line{}, err
	}

	return card, nil
}

func (c *Cards) MarkUnknown(card_id string) error {
	res, err := c.db.Exec(`UPDATE cards SET know = 0 WHERE id = $1`, card_id)
	if err != nil {
		return err
	}
	rowsUpdate, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Print(rowsUpdate)
	return nil
}
