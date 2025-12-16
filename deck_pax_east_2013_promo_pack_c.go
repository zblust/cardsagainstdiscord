package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-east-2013-promo-pack-c",
		Description: "PAX East 2013 Promo Pack C pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `I don't know exactly how I got the PAX plague, but I suspect it had something to do with %s.`},
			&PromptCard{Prompt: `Priests think about %s in their free time.`},
		},

		Responses: []ResponseCard{
			`Achieving a gangbang life goal`,
			`Charisma`,
			`Forgetting to move the damn Elf on the Shelf`,
			`Judging elves by the color of their skin and not by the content of their character.`,
			`Smashing all the pottery in a Pottery Barn in search of rupees.`,
			`The Klobb.`,
			`Vespene gas.`,
			`Wil Wheaton crashing an actual spaceship.`,
		},
	}

	AddPack(pack)
}
