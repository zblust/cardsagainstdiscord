package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-prime-2014---panel-cards",
		Description: "PAX Prime 2014 - Panel Cards pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Bob Ross's little-known first show was called "The Joy of %s."`},
			&PromptCard{Prompt: `Buzzfeed presents: 10 pictures of %s that look like %s.`},
			&PromptCard{Prompt: `During my first game of D&D, I accidentally summoned %s.`},
			&PromptCard{Prompt: `Like %s, State Farm is there.`},
			&PromptCard{Prompt: `The Discovery Channel presents: %s week.`},
		},

		Responses: []ResponseCard{
			`A neck beard that is 10% cheese.`,
			`No survivors.`,
			`penetrable stuff`,
			`Pooping as quietly as possible.`,
			`The beautiful sport of Turkish oil wrestling.`,
		},
	}

	AddPack(pack)
}
