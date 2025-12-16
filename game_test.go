package cardsagainstdiscord

import (
	"testing"
)

func TestNextCardCzar(t *testing.T) {
	players := []*Player{
		{ID: 1, Playing: true, InGame: true},
		{ID: 5, Playing: true, InGame: true},
		{ID: 2, Playing: true, InGame: true},
	}

	current := NextCardCzar(players, 0)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}

	current = NextCardCzar(players, current)
	if current != 2 {
		t.Error("Got ", current, " exected 2")
	}

	current = NextCardCzar(players, current)
	if current != 5 {
		t.Error("Got ", current, " exected 5")
	}

	current = NextCardCzar(players, current)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}
}

func TestNextCardCzar2(t *testing.T) {
	players := []*Player{
		{ID: 5, Playing: true, InGame: true},
		{ID: 1, Playing: true, InGame: true},
		{ID: 2, Playing: true, InGame: true},
	}

	current := NextCardCzar(players, 0)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}

	current = NextCardCzar(players, current)
	if current != 2 {
		t.Error("Got ", current, " exected 2")
	}

	current = NextCardCzar(players, current)
	if current != 5 {
		t.Error("Got ", current, " exected 5")
	}

	current = NextCardCzar(players, current)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}
}

func TestNextCardCzar3(t *testing.T) {
	players := []*Player{
		{ID: 5, Playing: true, InGame: true},
		{ID: 1, Playing: true, InGame: true},
		{ID: 2, Playing: true, InGame: true},
		{ID: 3, Playing: true, InGame: true},
	}

	current := NextCardCzar(players, 0)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}

	current = NextCardCzar(players, current)
	if current != 2 {
		t.Error("Got ", current, " exected 2")
	}

	current = NextCardCzar(players, current)
	if current != 3 {
		t.Error("Got ", current, " exected 3")
	}

	current = NextCardCzar(players, current)
	if current != 5 {
		t.Error("Got ", current, " exected 5")
	}

	current = NextCardCzar(players, current)
	if current != 1 {
		t.Error("Got ", current, " exected 1")
	}
}

func TestDupeResponses(t *testing.T) {
	for k, pack := range Packs {
		g := &Game{
			Packs: []string{k},
		}

		pickedResponses := make([]ResponseCard, 0, len(pack.Responses))
		for i := 0; i < len(pack.Responses); i++ {
			card := g.getRandomResponseCard()
			if card == BlankCard {
				i--
				continue
			}

			for _, v := range pickedResponses {
				if v == card {
					t.Error(k, ": Got duplicate response: ", v)
					break
				}
			}

			pickedResponses = append(pickedResponses, card)
		}
	}
}

func TestDupePrompts(t *testing.T) {
	for k, pack := range Packs {
		g := &Game{
			Packs: []string{k},
		}

		pickedPrompts := make([]string, 0, len(pack.Prompts))
		for i := 0; i < len(pack.Prompts); i++ {
			prompt := g.randomPrompt()

			for _, v := range pickedPrompts {
				if v == prompt.Prompt {
					t.Error(k, ": Got duplicate prompt: ", v)
					break
				}
			}

			pickedPrompts = append(pickedPrompts, prompt.Prompt)
		}
	}
}

func TestSameCardAgainPlaceholder(t *testing.T) {
	prompt := &PromptCard{
		Prompt:  "You want %s? You can't handle %0!",
		NumPick: 1,
	}

	placeholder := prompt.PlaceHolder()
	expected := "You want \\_\\_\\_\\_\\_? You can't handle [SAME CARD AGAIN]!"
	if placeholder != expected {
		t.Errorf("PlaceHolder() = %q, expected %q", placeholder, expected)
	}
}

func TestSameCardAgainWithCards(t *testing.T) {
	prompt := &PromptCard{
		Prompt:  "You want %s? You can't handle %0!",
		NumPick: 1,
	}

	cards := []ResponseCard{"the truth"}
	result := prompt.WithCards(cards)
	expected := "You want **the truth**? You can't handle **the truth**!"
	if result != expected {
		t.Errorf("WithCards() = %q, expected %q", result, expected)
	}
}

func TestSameCardAgainWithCardsMultiple(t *testing.T) {
	prompt := &PromptCard{
		Prompt:  "Let's take it from the top, and remember, you are %s. Show me %0.",
		NumPick: 1,
	}

	cards := []ResponseCard{"a dinosaur"}
	result := prompt.WithCards(cards)
	expected := "Let's take it from the top, and remember, you are **a dinosaur**. Show me **a dinosaur**."
	if result != expected {
		t.Errorf("WithCards() = %q, expected %q", result, expected)
	}
}

func TestSameCardAgainMultipleOccurrences(t *testing.T) {
	// Test a prompt with %0 appearing multiple times
	prompt := &PromptCard{
		Prompt:  "I love %s! %0 is the best! Give me more %0!",
		NumPick: 1,
	}

	cards := []ResponseCard{"pizza"}
	result := prompt.WithCards(cards)
	expected := "I love **pizza**! **pizza** is the best! Give me more **pizza**!"
	if result != expected {
		t.Errorf("WithCards() = %q, expected %q", result, expected)
	}
}

func TestSameCardAgainWithMultipleCards(t *testing.T) {
	// Test a prompt with two picks and references to both
	prompt := &PromptCard{
		Prompt:  "First there was %s, then came %s, but I prefer %0 over %1.",
		NumPick: 2,
	}

	cards := []ResponseCard{"fire", "ice"}
	result := prompt.WithCards(cards)
	expected := "First there was **fire**, then came **ice**, but I prefer **fire** over **ice**."
	if result != expected {
		t.Errorf("WithCards() = %q, expected %q", result, expected)
	}
}

func TestSameCardAgainInPacks(t *testing.T) {
	// Verify that the theatre pack prompt has correct NumPick
	pack := Packs["theatre"]
	if pack == nil {
		t.Fatal("theatre pack not found")
	}

	var found bool
	for _, prompt := range pack.Prompts {
		if prompt.Prompt == "Let's take it from the top, and remember, you are %s. Show me %0." {
			found = true
			if prompt.NumPick != 1 {
				t.Errorf("Theatre pack prompt has NumPick = %d, expected 1", prompt.NumPick)
			}
		}
	}

	if !found {
		t.Error("Theatre pack prompt with %0 reference not found")
	}

	// Verify that the hidden gems pack prompt has correct NumPick
	pack = Packs["hidden-gems-bundle-a-few-new-cards-we-crammed-into-this-bundle-pack-amazon-exclusive"]
	if pack == nil {
		t.Fatal("hidden gems pack not found")
	}

	found = false
	for _, prompt := range pack.Prompts {
		if prompt.Prompt == "You want %s? You can't handle %0!" {
			found = true
			if prompt.NumPick != 1 {
				t.Errorf("Hidden gems pack prompt has NumPick = %d, expected 1", prompt.NumPick)
			}
		}
	}

	if !found {
		t.Error("Hidden gems pack prompt with %0 reference not found")
	}
}
