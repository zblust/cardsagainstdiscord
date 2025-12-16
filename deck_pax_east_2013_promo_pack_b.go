package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-east-2013-promo-pack-b",
		Description: "PAX East 2013 Promo Pack B pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Action stations! Action stations! Set condition one throughout the fleet and brace for %s!`},
			&PromptCard{Prompt: `In the final round of this year's Omegathon, Omeganauts must face off in a game of %s.`},
		},
		Responses: []ResponseCard{
			`Getting inside the Horadric Cube with a hot babe and pressing the transmute button.`,
			`Loan sharks or left sharks, whichever`,
			`Punching a tree to gather wood.`,
			`Sharpie lipstick`,
			`Spending the year's insulin budget on Warhammer 40k figurines.`,
			`The desparate girl who wanted to be a mom so bad that she robbed a sperm bank.`,
			`The rocket launcher.`,
			`Violating the First Law of Robotics.`,
		},
	}
	AddPack(pack)
}
