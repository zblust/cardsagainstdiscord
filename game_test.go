package cardsagainstdiscord

import (
	"testing"
)

func TestNextCardCzar(t *testing.T) {
	players := []*Player{
		{ID: "100000000000000001", Playing: true, InGame: true},
		{ID: "100000000000000005", Playing: true, InGame: true},
		{ID: "100000000000000002", Playing: true, InGame: true},
	}

	current := NextCardCzar(players, "")
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000002" {
		t.Error("Got ", current, " expected 100000000000000002")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000005" {
		t.Error("Got ", current, " expected 100000000000000005")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
	}
}

func TestNextCardCzar2(t *testing.T) {
	players := []*Player{
		{ID: "100000000000000005", Playing: true, InGame: true},
		{ID: "100000000000000001", Playing: true, InGame: true},
		{ID: "100000000000000002", Playing: true, InGame: true},
	}

	current := NextCardCzar(players, "")
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000002" {
		t.Error("Got ", current, " expected 100000000000000002")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000005" {
		t.Error("Got ", current, " expected 100000000000000005")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
	}
}

func TestNextCardCzar3(t *testing.T) {
	players := []*Player{
		{ID: "100000000000000005", Playing: true, InGame: true},
		{ID: "100000000000000001", Playing: true, InGame: true},
		{ID: "100000000000000002", Playing: true, InGame: true},
		{ID: "100000000000000003", Playing: true, InGame: true},
	}

	current := NextCardCzar(players, "")
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000002" {
		t.Error("Got ", current, " expected 100000000000000002")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000003" {
		t.Error("Got ", current, " expected 100000000000000003")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000005" {
		t.Error("Got ", current, " expected 100000000000000005")
	}

	current = NextCardCzar(players, current)
	if current != "100000000000000001" {
		t.Error("Got ", current, " expected 100000000000000001")
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
	expected := "You want \\_\\_\\_\\_\\_? You can't handle [FIRST CARD AGAIN]!"
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

	// Test PlaceHolder shows [FIRST CARD AGAIN]
	placeholder := prompt.PlaceHolder()
	expectedPlaceholder := "I love \\_\\_\\_\\_\\_! [FIRST CARD AGAIN] is the best! Give me more [FIRST CARD AGAIN]!"
	if placeholder != expectedPlaceholder {
		t.Errorf("PlaceHolder() = %q, expected %q", placeholder, expectedPlaceholder)
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

	// Test PlaceHolder shows different labels for different positions
	placeholder := prompt.PlaceHolder()
	expectedPlaceholder := "First there was \\_\\_\\_\\_\\_, then came \\_\\_\\_\\_\\_, but I prefer [FIRST CARD AGAIN] over [SECOND CARD AGAIN]."
	if placeholder != expectedPlaceholder {
		t.Errorf("PlaceHolder() = %q, expected %q", placeholder, expectedPlaceholder)
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

func TestPackBlacklist(t *testing.T) {
	// Test case 1: All packs except first and bluebox
	selectedPacks, err := ProcessPacks("*", "-first", "-bluebox")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify first and bluebox are not in the list
	for _, pack := range selectedPacks {
		if pack == "first" || pack == "bluebox" {
			t.Errorf("Blacklisted pack %s should not be in selected packs", pack)
		}
	}

	// Verify at least some packs are selected
	if len(selectedPacks) < 1 {
		t.Error("No packs selected when using blacklist")
	}
	
	// Verify we have fewer packs than total (since we excluded 2)
	if len(selectedPacks) >= len(Packs) {
		t.Error("Should have fewer packs when using blacklist")
	}

	// Test case 2: Whitelist should still work
	whitelistPacks, err := ProcessPacks("first", "third")
	if err != nil {
		t.Errorf("Unexpected error for whitelist: %v", err)
	}

	if len(whitelistPacks) != 2 {
		t.Errorf("Expected 2 packs in whitelist, got %d", len(whitelistPacks))
	}
	
	// Verify the packs are correct
	packMap := make(map[string]bool)
	for _, pack := range whitelistPacks {
		packMap[pack] = true
	}
	if !packMap["first"] || !packMap["third"] {
		t.Error("Whitelist should contain first and third")
	}
}

func TestPackBlacklistUnknownPack(t *testing.T) {
	// Test: Blacklist unknown pack should return error
	_, err := ProcessPacks("*", "-unknownpack")
	if err == nil {
		t.Error("Expected error when blacklisting unknown pack")
	}

	// Verify it's the right error type
	if _, ok := err.(*ErrUnknownPack); !ok {
		t.Errorf("Expected ErrUnknownPack, got %T", err)
	}
}

func TestPackBlacklistOnly(t *testing.T) {
	// Test: Only blacklist without * should return error
	_, err := ProcessPacks("-first", "-bluebox")
	if err != ErrNoPacks {
		t.Errorf("Expected ErrNoPacks when only blacklist is provided, got %v", err)
	}
}

func TestPackAllBlacklisted(t *testing.T) {
	// Get all pack names
	allPackNames := make([]string, 0, len(Packs))
	allPackNames = append(allPackNames, "*")
	for k := range Packs {
		allPackNames = append(allPackNames, "-"+k)
	}

	// Test: Blacklist all packs should return error
	_, err := ProcessPacks(allPackNames...)
	if err != ErrNoPacks {
		t.Errorf("Expected ErrNoPacks when all packs are blacklisted, got %v", err)
	}
}
