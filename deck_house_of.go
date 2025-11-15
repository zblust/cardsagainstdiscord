package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "House of Cards Pack",
		Description: "House of Cards Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `A wise man said, "Everything is about sex. Except sex. Sex is about _."`},
			&PromptCard{Prompt: `Because you enjoyed _, we thought you'd like _.`},
			&PromptCard{Prompt: `Cancel all my meetings. We've got a situation with _ that requires my immediate attention.`},
			&PromptCard{Prompt: `Corruption. Betrayal. _. Coming soon to Netflix, "House of _."`},
			&PromptCard{Prompt: `I can't believe Netflix is using _ to promote House of Cards.`},
			&PromptCard{Prompt: `I'm not going to lie. I despise _. There, I said it.`},
			&PromptCard{Prompt: `If you need him to, Remy Danton can pull some strings and get you _, but it'll cost you.`},
			&PromptCard{Prompt: `Our relationship is strictly professional. Let's not complicate things with _.`},
			&PromptCard{Prompt: `We're not like other news organizations. Here at Slugline, we welcome _ in the office.`},
		},

		Responses: []ResponseCard{
			`25 shitty jokes about House of Cards.`,
			`A childless marriage.`,
			`A homoerotic subplot.`,
			`A mucous plug.`,
			`An older woman who knows her way around the penis.`,
			`An origami swan that's some kind of symbol?`,
			`Carbon monoxide poisoning.`,
			`Discharging a firearm in a residential area.`,
			`Forcing a handjob on a dying man.`,
			`Getting eaten out while on the phone with Dad.`,
			`Making it look like a suicide.`,
			`My constituents.`,
			`Punching a guy through a wall.`,
			`Ribs so good they transcend race and class.`,
			`Strangling a dog to make a point to the audience.`,
			`The sensitive European photographer who's fucking my wife.`,
		},
	}

	AddPack(pack)
}
