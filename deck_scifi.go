package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "sci-fi",
		Description: "Sci-Fi Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Computer! Display %s on screen. Enhance.`},
			&PromptCard{Prompt: `Feeling nostalgic when you remember %s from your childhood.`},
			&PromptCard{Prompt: `Madam President, the asteroid is headed directly for Earth and there's only one thing that can stop it: %s.`},
			&PromptCard{Prompt: `This won't be like negotiating with the Vogons. Humans only respond to one thing: %s.`},
			&PromptCard{Prompt: `What is the answer to life, the universe, and everything?`},
			&PromptCard{Prompt: `You have violated the Prime Directive! You exposed an alien culture to %s before they were ready.`},
			&PromptCard{Prompt: `You're not going to believe this, but I'm you from the future! You've got to stop %s.`},
		},

		Responses: []ResponseCard{
			`[A picture of Sean Connery in the movie Zardoz]`,
			`A hazmat suit full of farts.`,
			`A misty room full of glistening egg sacs.`,
			`A planet-devouring space worm named Rachel.`,
			`A protagonist with no qualities.`,
			`An alternate history where Hitler was gay but he still killed all those people.`,
			`Beep beep boop beep boop.`,
			`Cheerful blowjob robots.`,
			`Cosmic bowling.`,
			`Darmok and Jalad at Tanagra.`,
			`Frantically writing equations on a chalkboard.`,
			`Funkified aliens from the planet Groovius.`,
			`Going too far with science and bad things happening.`,
			`How great of a movie Men In Black was.`,
			`Laying thousands of eggs in a man's colon.`,
			`Masturbating Yoda's leathery turtle-penis.`,
			`Nine seasons of sexual tension with David Duchovny.`,
			`That girl from the Hunger Games.`,
			`The dystopia we're living in right now.`,
			`The ending of Lost.`,
			`Three boobs.`,
			`Trimming the poop out of Chewbacca's butt hair.`,
			`Vulcan sex-madness.`,
		},
	}

	AddPack(pack)
}
