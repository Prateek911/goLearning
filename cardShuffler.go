package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CardShuffler struct {
	cards []string
}

func NewCardShuffler(cards []string) *CardShuffler {
	return &CardShuffler{
		cards: cards,
	}
}

func (cs *CardShuffler) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	rng.Shuffle(len(cs.cards), func(i, j int) {
		cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i]
	})
}

func (cs *CardShuffler) DrawCard() (string, error) {
	if len(cs.cards) == 0 {
		return "", fmt.Errorf("no Cards left")
	}
	card := cs.cards[0]
	cs.cards = cs.cards[1:]
	return card, nil
}

func (cs *CardShuffler) DisplayCards() {
	fmt.Println("Remaining cards in the Deck:")
	for _, card := range cs.cards {
		fmt.Println(card)
	}
}

func main() {
	deckOfCards := []string{"Ace of Spades", "King of Hearts", "Queen of Diamonds", "Jack of Clubs"}
	shuffler := NewCardShuffler(deckOfCards)
	shuffler.Shuffle()

	var card string
	var err error

	if card, err = shuffler.DrawCard(); err == nil {
		fmt.Println("Drawn Card :", card)
		shuffler.DisplayCards()
	} else {
		fmt.Println(err)
	}

}
