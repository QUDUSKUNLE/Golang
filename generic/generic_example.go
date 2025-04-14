package generic

import (
	"math/rand"
	"time"
)

type PlayingCard struct {
	Suit  string
	Rank  string
}

func NewPlayingCard(suit, rank string) *PlayingCard {
	return &PlayingCard{
		Suit: suit,
		Rank: rank,
	}
}

func (c *PlayingCard) String() string {
	return c.Rank + " of " + c.Suit
}

type Deck[G any] struct {
	Cards []G
}

func (d *Deck[G]) AddCard(card G) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck[G]) RandomCard() G {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return d.Cards[r.Intn(len(d.Cards))]
}
