package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-prime-2015-food-pack-a-mango",
		Description: "PAX Prime 2015 Food Pack A (Mango)",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `I'm Bobby Flay, and if you can't stand %s, get out of the kitchen!`},
			&PromptCard{Prompt: `It's not delivery.
It's %s.`},
		},

		Responses: []ResponseCard{
			`A soccer ball to the crotch.`,
			`Golden Girls re-runs.`,
			`Kale.`,
			`Licking the flavor off of Doritos so you can reuse them as tortilla chips for your salsa.`,
			`Real News.`,
			`Switching bodies with mom for a day.`,
			`The Diary of Anne Spank: An XXX Parody.`,
			`What we think is meatloaf`,
		},
	}

	AddPack(pack)
}
