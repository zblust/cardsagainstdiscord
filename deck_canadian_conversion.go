package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "CAH: Canadian Conversion Kit",
		Description: "CAH: Canadian Conversion Kit pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Air Force fuckery is getting so bad that the brass are starting to blame _ on the lack of discipline.`},
			&PromptCard{Prompt: `CTV presents _, the story of _.`},
			&PromptCard{Prompt: `In an attempt to reach a wider audience, the Royal Ontario Museum has opened an interactive exhibit on _.`},
			&PromptCard{Prompt: `Next season on Celebrity Apprentice: _.`},
			&PromptCard{Prompt: `O say, does _ yet wave

o'er the land of the fre and the home of the brave?`},
			&PromptCard{Prompt: `The Royal Ontario Museum has just opened an interactive exhibit on _.`},
			&PromptCard{Prompt: `What's the crustiest?`},
			&PromptCard{Prompt: `When I am Prime Minister of Canada, I will create the Ministry of _.`},
		},

		Responses: []ResponseCard{
			`A Molson muscle.`,
			`A time travel paradox.`,
			`An identity crisis`,
			`Being catfished.`,
			`Burning Flipside.`,
			`Canada: America's hat.`,
			`Don Cherrys wardrobe.`,
			`hermaphroditical Italian pictures`,
			`Homoerotic ass slapping amongst athletes.`,
			`Mr. Froto's ring`,
			`Naked people.`,
			`Newly incorporated kamikaze tactics.`,
			`Poverty porn`,
			`Rob Ford`,
			`School friends' uneducated views on sex`,
			`Snow falling gently on the frozen body of an orphan boy.`,
			`Stephen Hawking`,
			`Systemic racism.`,
			`Test driving a used sex toy.`,
			`The Famous Five.`,
			`The FLQ`,
			`The Ohakune Carrot`,
			`The royal penis.`,
		},
	}

	AddPack(pack)
}
