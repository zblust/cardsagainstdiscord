package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "World Wide Web Pack",
		Description: "World Wide Web Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Don't worry, Penny! Go Go Gadget _!`},
			&PromptCard{Prompt: `I need you like _ needs _.`},
			&PromptCard{Prompt: `I'm just gonna stay in tonight. You know, Netflix and _.`},
			&PromptCard{Prompt: `Nothing says "I love you" like _.`},
			&PromptCard{Prompt: `Such _. Very _. Wow.`},
			&PromptCard{Prompt: `This app is basically Tinder, but for _.`},
			&PromptCard{Prompt: `TRIGGER WARNING: _.`},
			&PromptCard{Prompt: `What did I nickname my genitals?`},
			&PromptCard{Prompt: `You guys, you can buy _ on the dark web.`},
		},

		Responses: []ResponseCard{
			`A complete inability to understand anyone else's perspective.`,
			`A fun, sexy time at the nude beach.`,
			`A man from Craigslist.`,
			`A night of Taco Bell and anal sex.`,
			`A respectful discussion of race and gender on the Internet.`,
			`Cat massage.`,
			`Destroying Dick Cheney's last horcrux.`,
			`Game of Thrones spoilers.`,
			`Getting teabagged by a fifth grader in Call of Duty.`,
			`Goats screaming like people.`,
			`Googling.`,
			`Internet porn analysis paralysis.`,
			`Matching with Mom on Tinder.`,
			`My browser history.`,
			`My privileged white penis.`,
			`Pretending to be black.`,
			`Smash Mouth.`,
			`Taking a shit while running at full speed.`,
			`That thing on the Internet everyone's talking about.`,
			`Three years of semen in a shoebox.`,
			`Youtube comments.`,
		},
	}

	AddPack(pack)
}
