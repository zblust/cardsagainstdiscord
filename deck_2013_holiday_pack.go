package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "2013-holiday-pack",
		Description: "2013 Holiday Pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Because they are forbidden from masturbating, Mormons channel their repressed sexual energy into %s.`},
			&PromptCard{Prompt: `Blessed are you, Lord our God, creator of the universe, who has granted us %s.`},
			&PromptCard{Prompt: `But wait, there's more! If you order %s in the next 15 minutes, we'll throw in %s absolutely free!`},
			&PromptCard{Prompt: `GREETINGS HUMANS
I AM %s BOT
EXECUTING PROGRAM.`},
			&PromptCard{Prompt: `Heroin: a proud supporter of %s.`},
			&PromptCard{Prompt: `I really hope my grandma doesn't ask me to explain %s again.`},
			&PromptCard{Prompt: `Kids these days with their iPods and their Internet. In my day, all we needed to pass the time was %s.`},
			&PromptCard{Prompt: `Revealed: Why He Really Resigned! Pope Benedict's Secret Struggle with %s.`},
			&PromptCard{Prompt: `What's the only thing sexier than confidence?`},
		},
		Responses: []ResponseCard{
			`A magical tablet containing a world of unlimited pornography.`,
			`A simultaneous nightmare and wet dream starring Sigourney Weaver.`,
			`Being blind and deaf and having no limbs.`,
			`Breeding elves for their priceless semen.`,
			`Congress's flaccid penises withering away beneath their suit pants.`,
			`Finding out Santa isn't real.`,
			`Giving money and personal information to strangers on the Internet.`,
			`Having a strong opinion about Obamacare.`,
			`Jizzing into Santa's beard.`,
			`Making up for 10 years of shitty parenting with a PlayStation.`,
			`Moses gargling Jesus's balls while Shiva and the Buddha penetrate his divine hand holes.`,
			`People with cake in their mouths talking about how good cake is.`,
			`Piece of shit Christmas cards with no money in them.`,
			`Rudolph's bright red balls.`,
			`Slicing a ham in icy silence.`,
			`Swapping bodies with mom for a day.`,
			`The Grinch's musty, cum-stained pelt.`,
			`The Hawaiian goddess Kapo and her flying detachable vagina.`,
			`The royal afterbirth.`,
			`The shittier, Jewish version of Christmas.`,
			`These low, low prices!`,
		},
	}

	AddPack(pack)
}
