package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "TableTop Pack",
		Description: "TableTop Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Backers who supported Tabletop at the $25,000 level were astonished to receive %s from Wil Wheaton himself.`},
			&PromptCard{Prompt: `For my turn, I will spend four gold and allocate all three workers to %s.`},
			&PromptCard{Prompt: `Hey, you guys want to try this awesome new game? It's called %s.`},
		},

		Responses: []ResponseCard{
			`A disappointing season of Tabletop that's just about tables.`,
			`A German-style board game where you invade Poland.`,
			`A marriage-destroying game of The Resistance.`,
			`A Wesley Crusher blow-up doll.`,
			`A zombie with a tragic backstory.`,
			`An owlbear.`,
			`Condensing centuries of economic exploitation into 90 minutes of gaming fun.`,
			`SIX GOD DAMN HOURS OF FUCKING DIPLOMACY.`,
			`Spending 8 years in the Himalayas becoming a master of dice-rolling and resource allocation.`,
			`The pooping position.`,
			`The porn set that Tabletop is filmed on.`,
			`Victory points.`,
		},
	}

	AddPack(pack)
}
