package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX Prime 2014 - Panel Cards",
		Description: "PAX Prime 2014 - Panel Cards pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Bob Ross's little-known first show was called "The Joy of _."`},
			&PromptCard{Prompt: `Buzzfeed presents: 10 pictures of _ that look like _.`},
			&PromptCard{Prompt: `During my first game of D&D, I accidentally summoned _.`},
			&PromptCard{Prompt: `Like _, State Farm is there.`},
			&PromptCard{Prompt: `The Discovery Channel presents: _ week.`},
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
