package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "90s-nostalgia-pack",
		Description: "90s Nostalgia Pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Believe it or not Jim Carrey can do a dead-on impression of %s.`},
			&PromptCard{Prompt: `How did Stella get her groove back? %s`},
			&PromptCard{Prompt: `I'm a bitch, I'm a lover, I'm a child, I'm %s.`},
			&PromptCard{Prompt: `It's Morphin' Time! Mastodon! Pterodactyl! Triceratops! Sabertooth Tiger! %s!`},
			&PromptCard{Prompt: `Siskel and Ebert have panned %s as "poorly conceived" and "sloppily executed."`},
			&PromptCard{Prompt: `Tonight on SNICK: "Are You Afriad of %s?`},
			&PromptCard{Prompt: `Up next on Nickelodeon: "Clarissa Explains %s."`},
		},
		Responses: []ResponseCard{
			`A bus that will explode if it goes under 50 miles per hour.`,
			`A mulatto, an albino, a mosquito, and my libido.`,
			`A threesome with 1996 Denise Richards and 1999 Denise Richards.`,
			`Angels interfering in an otherwise fair baseball game.`,
			`Cool 90s up-in-the-front hair.`,
			`Deregulating the mortgage market.`,
			`Freeing Willy.`,
			`Getting caught up in the CROSSFIRE™.`,
			`Jerking off to a 10-second RealMedia clip.`,
			`Kurt Cobain's death.`,
			`Liking big butts and not being able to lie about it.`,
			`Log™.`,
			`Painting with all the colors of the wind.`,
			`Pamela Anderson's boobs running in slow motion.`,
			`Patti Mayonnaise.`,
			`Pizza in the morning, pizza in the evening, pizza at supper time.`,
			`Pure Moods, Vol. 1.`,
			`Several Michael Keatons`,
			`Stabbing the shit out of a Capri Sun.`,
			`Sucking the president's dick.`,
			`Sunny D! Alright!`,
			`The Great Cornholio.`,
			`The Y2K bug.`,
			`Wearing Nicolas Cage's face.`,
			`Yelling "girl power!" and doing a high kick.`,
		},
	}

	AddPack(pack)
}
