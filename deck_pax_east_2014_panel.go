package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX East 2014 - Panel Cards",
		Description: "PAX East 2014 - Panel Cards pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `_ was totally worth the trauma.`},
			&PromptCard{Prompt: `Let me tell you about my new startup. It's basically _, but for _.`},
		},

		Responses: []ResponseCard{
			`A floor that is literally made of lava.`,
			`All this liquid in my mouth.`,
			`Exciting content!`,
			`Giving a dolphin a handjob for science.`,
			`Hoes in different area codes.`,
			`Rubbing chocolate pudding all over Bill Cosby's nipples.`,
			`Stepping on a god damn friggin' LEGO.`,
			`What The Rock was really cooking.`,
		},
	}

	AddPack(pack)
}
