package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX Prime 2015 Food Pack C (Cherry)",
		Description: "PAX Prime 2015 Food Pack C (Cherry) pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Excuse me, waiter. Could take this back? This soup tastes like _.`},
			&PromptCard{Prompt: `Now on Netflix: Jiro Dreams of _.`},
		},

		Responses: []ResponseCard{
			`A Mexican child trapped inside of a burrito.`,
			`Clams Attempt Harmonica`,
			`Committing suicide.`,
			`Faulty UIDs.`,
			`Jizz-flavored coffee.`,
			`The Helvetica Scenario.`,
			`the hot duke`,
			`the incestuous pleasure of his bed`,
		},
	}

	AddPack(pack)
}
