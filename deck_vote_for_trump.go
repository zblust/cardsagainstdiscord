package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "vote-for-trump",
		Description: "Vote For Trump Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `According to Arizona's stand-your-ground law, you're allowed to shoot someone if they're %s.`},
			&PromptCard{Prompt: `It's 3AM.  The red phone rings.  It's %s.  Who do you want answering?`},
			&PromptCard{Prompt: `Trump's great!  Trump's got %s.  I love that.`},
		},

		Responses: []ResponseCard{
			`A liberal bias.`,
			`Actually voting for Donald Trump to be President of the actual United States.`,
			`Conservative talking points.`,
			`Courageousely going ahead with that racist comment.`,
			`Dispelling this fiction that Barack Obama doesn't know what he's doing.`,
			`Fully appreciating naked Morris dancing.`,
			`Growing up and becoming a Republican.`,
			`Hating Hilary Clinton.`,
			`Jeb!`,
			`Shouting the loudest.`,
			`Sound fiscal policy.`,
			`The good, hardworking people of Dubuque, Iowa.`,
		},
	}

	AddPack(pack)
}
