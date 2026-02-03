package cardsagainstdiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
)

var Packs = make(map[string]*CardPack)

func AddPack(pack *CardPack) {
	// Count picks
	for _, v := range pack.Prompts {
		numPicks := strings.Count(v.Prompt, "%s")
		if numPicks == 0 {
			v.Prompt += " %s"
			v.NumPick = 1
		} else {
			// %0, %1, etc. are references to already-picked cards, not new picks
			// So we only count %s placeholders as actual picks needed
			v.NumPick = numPicks
		}
	}

	Packs[pack.Name] = pack
}

type CardPack struct {
	Name        string
	Description string
	Prompts     []*PromptCard
	Responses   []ResponseCard
}

type PromptCard struct {
	Prompt  string
	NumPick int
}

var (
	EscaperReplacer = strings.NewReplacer("*", "\\*", "_", "\\_")
)

const (
	maxCardReferences = 10 // Maximum number of card position references (e.g., %0 through %9)
)

func (p *PromptCard) PlaceHolder() string {
	s := strings.Replace(p.Prompt, "%s", "_____", -1)
	// Replace %0, %1, etc. with [FIRST CARD AGAIN], [SECOND CARD AGAIN], etc.
	cardOrdinals := []string{"FIRST", "SECOND", "THIRD", "FOURTH", "FIFTH", "SIXTH", "SEVENTH", "EIGHTH", "NINTH", "TENTH"}
	for i := 0; i < maxCardReferences && i < len(cardOrdinals); i++ {
		placeholder := fmt.Sprintf("%%%d", i)
		replacement := fmt.Sprintf("[%s CARD AGAIN]", cardOrdinals[i])
		s = strings.Replace(s, placeholder, replacement, -1)
	}
	s = strings.Replace(s, "%%", `%`, -1)

	s = EscaperReplacer.Replace(s)

	return s
}

func (p *PromptCard) WithCards(cards interface{}) string {
	args := make([]interface{}, p.NumPick)
	switch t := cards.(type) {
	case []string:
		for i, v := range t {
			args[i] = "**" + v + "**"
		}
	case []ResponseCard:
		for i, v := range t {
			args[i] = "**" + v + "**"
		}
	}

	// Replace %0, %1, etc. with %s and collect duplicates to append
	// For example: "I love %s! %0 is great!" becomes "I love %s! %s is great!"
	// We need args[0] twice: once for %s and once for %0
	s := p.Prompt
	duplicates := make([]interface{}, 0)
	for i := 0; i < p.NumPick && i < maxCardReferences; i++ {
		placeholder := fmt.Sprintf("%%%d", i)
		count := strings.Count(s, placeholder)
		if count > 0 {
			s = strings.Replace(s, placeholder, "%s", -1)
			// Each %0 reference needs a copy of args[i] for fmt.Sprintf
			for j := 0; j < count; j++ {
				duplicates = append(duplicates, args[i])
			}
		}
	}

	// Append all duplicates at once
	args = append(args, duplicates...)

	s = fmt.Sprintf(s, args...)
	// s = EscaperReplacer.Replace(s)
	return s
}

type ResponseCard string

type SessionProvider interface {
	SessionForGuild(guildID string) *discordgo.Session
}

type StaticSessionProvider struct {
	Session *discordgo.Session
}

func (sp *StaticSessionProvider) SessionForGuild(guildID string) *discordgo.Session {
	return sp.Session
}

var (
	ErrGameAlreadyInChannel = errors.New("Already a active game in this channel")
	ErrPlayerAlreadyInGame  = errors.New("Player already in a game")
	ErrGameNotFound         = errors.New("Game not found")
	ErrGameFull             = errors.New("Game is full")
	ErrNoPacks              = errors.New("No packs specified")
	ErrNotGM                = errors.New("You're not the game master")
	ErrStoppedAlready       = errors.New("Game already stopped")
	ErrPlayerNotInGame      = errors.New("Player not in your game")
)

type ErrUnknownPack struct {
	PassedPack string
}

func (e *ErrUnknownPack) Error() string {
	return "Unknown pack `" + e.PassedPack + "`"
}

func HumanizeError(err error) string {
	err = errors.Cause(err)

	if err == ErrGameAlreadyInChannel || err == ErrPlayerAlreadyInGame || err == ErrGameNotFound || err == ErrGameFull || err == ErrNoPacks || err == ErrNotGM || err == ErrStoppedAlready || err == ErrPlayerNotInGame {
		return err.Error()
	}

	if c, ok := err.(*ErrUnknownPack); ok {
		return c.Error()
	}

	return ""
}
