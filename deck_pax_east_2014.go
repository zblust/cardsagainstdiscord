package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX East 2014",
		Description: "PAX East 2014 pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `%s is way too much responsibility for me.`},
			&PromptCard{Prompt: `(insert name) died when %s.`},
			&PromptCard{Prompt: `Unfortunately, Neo, no one can be told what %s is. You have to see it for yourself.`},
			&PromptCard{Prompt: `What the hell?! They added a 6/6 with flying, trample and %s.`},
			&PromptCard{Prompt: `You think you have defeated me? Well, let's see how you handle %s.`},
		},

		Responses: []ResponseCard{
			`A giant mechanical bird with a tragic backstory.`,
			`A grumpy spiritual director.`,
			`All of the good times and premium gaming entertainment available to you in the Kickstarter room.`,
			`Attacking from Kamchatka.`,
			`Collecting all seven power crystals.`,
			`Demons and shit.`,
			`Endless pleasure.`,
			`Futuristic death sports.`,
			`Getting bitten by a radioactive spider and then battling leukemia for 30 years.`,
			`KHAAAAAAAAAN!`,
			`Lagging out.`,
			`Mistakenly hitting on a League of Legends statue.`,
			`Separation of merch and state.`,
			`Serana William's black sweaty scrotum`,
			`Stuffing my underwear with pancakes.`,
			`Taking 3 hours to go on a 15 minute errand`,
			`The imagination of Peter Jackson.`,
			`The old gods.`,
			`The pure, Zen-like state that exists between micro and macro.`,
			`The Star Wars universe.`,
			`Whatever Final Fantasy bullshit happened this year.`,
			`Xena, Warrior Princess.`,
		},
	}

	AddPack(pack)
}
