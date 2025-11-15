package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "jack-white-show-pack",
		Description: "Jack White Show Pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `When Jack White performed at the Aragon Ballroom in Chicago on Nov. 19, 2018, he surprised fans with %s.`},
		},
		Responses: []ResponseCard{
			`A bizarre partnership with Cards Against Humanity.`,
			`A quiet, poignant acoustic guitar ballad.`,
			`Jack White.`,
			`Jack White's collection of taxidermied animal heads.`,
			`Jack White's mom coming on stage and dancing.`,
			`Jack White's testosterone-fueled dementia.`,
		},
	}

	AddPack(pack)
}
