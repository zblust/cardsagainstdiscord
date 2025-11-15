package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Retail Mini Pack",
		Description: "Retail Mini Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `How are the writers of Cards Against Humanity spending your $25?`},
			&PromptCard{Prompt: `Looking to earn big bucks? Learn how to make _ work for you!`},
		},

		Responses: []ResponseCard{
			`A teenage boy gunning for a handjob.`,
			`Feeding a man a pie made of his own children.`,
			`Ironically buying a trucker hat and then ironically being a trucker for 38 years.`,
		},
	}

	AddPack(pack)
}
