package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-east-2014-panel-cards",
		Description: "PAX East 2014 - Panel Cards",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `%s was totally worth the trauma.`},
			&PromptCard{Prompt: `Let me tell you about my new startup. It's basically %s, but for %s.`},
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
