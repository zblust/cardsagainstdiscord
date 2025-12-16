package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "pax-prime-2015-food-pack-b-coconut",
		Description: "PAX Prime 2015 Food Pack B (Coconut)",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Aw babe, your burps smell like %s!`},
			&PromptCard{Prompt: `Don't miss Rachel Ray's hit new show, Cooking with %s.`},
		},

		Responses: []ResponseCard{
			`A Benedict Cumberbatch RealDoll™.`,
			`A joyless vegan patty.`,
			`A tablespoon of thick, custardy puss.`,
			`Being eskimo brothers with your father-in-law.`,
			`Kevin Hart`,
			`Not knowing when to shut up.`,
			`Soup that's better than pussy.`,
			`Sucking each other's penises for hours on end.`,
		},
	}

	AddPack(pack)
}
