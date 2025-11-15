package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX East 2013 Promo Pack A",
		Description: "PAX East 2013 Promo Pack A pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `I have an idea even better than Kickstarter, and it's called _starter.`},
			&PromptCard{Prompt: `You have been waylaid by _ and must defend yourself.`},
		},
		Responses: []ResponseCard{
		},
	}
	AddPack(pack)
}
