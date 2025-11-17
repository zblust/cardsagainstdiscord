package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Reject Pack 2",
		Description: "Reject Pack 2 pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `America is hungry. America needs %s.`},
			&PromptCard{Prompt: `Astronomers have discovered that the universe consists of 5% ordinary matter, 25% dark matter, and 70% %s.`},
			&PromptCard{Prompt: `BowWOW! is the first pet hotel in LA that offers %s for dogs.`},
			&PromptCard{Prompt: `Hey, whatever happened to Renee Zellweger?`},
			&PromptCard{Prompt: `Housekeeping! You want %s?`},
			&PromptCard{Prompt: `In bourgeois society, capital is independent and has individuality, while the living person is %s.`},
			&PromptCard{Prompt: `Some men aren't looking for anything logical, like money. They can't be bought, bullied, reasoned or negotiated with. Some men just want %s.`},
			&PromptCard{Prompt: `What's wrong with these gorillas?`},
			&PromptCard{Prompt: `Why did the fusion exchanges end?`},
			&PromptCard{Prompt: `You say tomato, I say %s.`},
		},

		Responses: []ResponseCard{
			`A dick so big and so black that not even light can escape its pull.`,
			`A double murder suicide barbeque.`,
			`A primordial soup and salad bar.`,
			`A stack of bunnies in a trenchcoat.`,
			`At least three ducks.`,
			`Becoming engorged with social justice jelly and secreting a thinkpiece.`,
			`Being the absolute worst.`,
			`Discolored Puss`,
			`Greg Kinnear's terrible lightning breath.`,
			`Mitt Romney's eight sons Kip, Sam, Trot, Fergis, Toolshed, Grisham, Hawkeye, and Thorp.`,
			`Mr. and Mrs. Tambourine Man's jingle-jangle morning sex.`,
			`Mushy tushy.`,
			`Ringo Starr & His All-Starr Band.`,
			`Sandy feet`,
			`Saving the Rainforest Cafe.`,
			`Sir Thomas More's Fruitopia™.`,
			`Sweating it out on the streets of a runaway American Dream.`,
			`That one leftover screw.`,
			`That thing politicians do with their thumbs when they talk.`,
			`The spooky skeleton under my skin.`,
			`The token minority.`,
			`These dolphins.`,
			`Three hairs from the silver-golden head of Galadriel.`,
			`Water so cold it turned into a rock.`,
		},
	}

	AddPack(pack)
}
