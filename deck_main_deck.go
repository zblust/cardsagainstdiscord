package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "cah-main-deck",
		Description: "CAH: Main Deck",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Channel 4 presents "%s: the Story of %s."`},
			&PromptCard{Prompt: `Dear Agony Aunt, I'm having some trouble with %s and would like your advice.`},
			&PromptCard{Prompt: `Hey guys, welcome to TGI Fridays! Would you like to start the night off right with %s?`},
			&PromptCard{Prompt: `Mate, do not go in that toilet. There's %s in there.`},
			&PromptCard{Prompt: `This season at the Old Vic, Samuel Beckett's classic existential play: Waiting for %s.`},
		},
		Responses: []ResponseCard{
			`A comprehensive understanding of the Irish backstop.`,
			`A deep-rooted fear of the working class.`,
			`A general lack of purpose.`,
			`A Ginsters pasty and three cans of Monsters Energy.`,
			`A meat raffle!`,
			`A slightly salty toad in the hole.`,
			`Ainsley Harriott.`,
			`All my gentleman suitors.`,
			`Barely making Â£15,000 a year.`,
			`Being a witch.`,
			`Being fucking stupid.`,
			`Blood, sweat, and tears.`,
			`Brutal austerity.`,
			`cunninglus`,
			`Danny Dyer.`,
			`Discovering he's a Tory.`,
			`Flat out not giving a shit.`,
			`Getting the same Boots Meal Deal every day for six years.`,
			`James fucking Cordon.`,
			`Licking the Queen.`,
			`Martin Lewis, Money Saving Expert.`,
			`My Uber driver, Ajay.`,
			`Nicki Minaj.`,
			`Ryanair.`,
			`Scottish independence.`,
			`Sitting in a jar of vinegar all night because I am gherkin.`,
			`Slapping your knees to signal your imminent departure.`,
			`The bastard seagull who stole my chips.`,
			`The Hair Plug Club`,
			`The petty troubles of the aristrocracy.`,
			`The Smell of a Primark.`,
			`The Strictly Come Dancing final.`,
			`The Welsh.`,
			`Waking up half-naked in a Wetherspoons car park.`,
		},
	}

	AddPack(pack)
}
