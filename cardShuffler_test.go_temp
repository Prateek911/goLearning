package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	// Initialize a CardShuffler with a known set of cards
	cards := []string{"Ace of Spades", "King of Hearts", "Queen of Diamonds", "Jack of Clubs"}
	shuffler := NewCardShuffler(cards)

	// Shuffle the cards
	shuffler.Shuffle()

	// Check if the length of the shuffled cards remains the same
	if len(shuffler.cards) != len(cards) {
		t.Errorf("Expected the length of shuffled cards (%d) to be the same as the original (%d)", len(shuffler.cards), len(cards))
	}

	// Check if the shuffled cards contain the same elements as the original set
	if !reflect.DeepEqual(shuffler.cards, cards) {
		t.Errorf("Expected shuffled cards to contain the same elements as the original")
	}
}

func TestDrawCard(t *testing.T) {
	// Initialize a CardShuffler with a known set of cards
	cards := []string{"Ace of Spades", "King of Hearts", "Queen of Diamonds", "Jack of Clubs"}
	shuffler := NewCardShuffler(cards)

	// Draw a card
	card, err := shuffler.DrawCard()

	// Check if drawing a card produced no error
	if err != nil {
		t.Errorf("Error while drawing a card: %v", err)
	}

	// Check if a card is drawn
	if card == "" {
		t.Errorf("Deck is empty : %v", err)
	}

}
