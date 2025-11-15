package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "house-of-cards-pack",
		Description: "House of Cards Pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `A wise man said, "Everything is about sex. Except sex. Sex is about %s."`},
			&PromptCard{Prompt: `Because you enjoyed %s, we thought you'd like %s.`},
			&PromptCard{Prompt: `Cancel all my meetings. We've got a situation with %s that requires my immediate attention.`},
			&PromptCard{Prompt: `Corruption. Betrayal. %s. Coming soon to Netflix, "House of %s."`},
			&PromptCard{Prompt: `I can't believe Netflix is using %s to promote House of Cards.`},
			&PromptCard{Prompt: `I'm not going to lie. I despise %s. There, I said it.`},
			&PromptCard{Prompt: `If you need him to, Remy Danton can pull some strings and get you %s, but it'll cost you.`},
			&PromptCard{Prompt: `Our relationship is strictly professional. Let's not complicate things with %s.`},
			&PromptCard{Prompt: `We're not like other news organizations. Here at Slugline, we welcome %s in the office.`},
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
