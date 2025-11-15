package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Dad Pack",
		Description: "Dad Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Coming up on Turner Classic Movies:
Sean Connery and Alec Baldwin star in "The Hunt for _."`},
			&PromptCard{Prompt: `FW: re:
FBI WARNING!!! VIOLENT GANGS USING _ TO RECRUIT CHILDREN!`},
			&PromptCard{Prompt: `Hey, kids.
I'm Connor's dad, but you can call me Mr. _.`},
			&PromptCard{Prompt: `So apparently Dad was searching Pornhub for "hot milf _."`},
			&PromptCard{Prompt: `Whaddya think, money grows on trees? I'm not paying for _!`},
			&PromptCard{Prompt: `Young lady, you better knock it off with _ or you're grounded!`},
		},

		Responses: []ResponseCard{
			`A positive male role model.`,
			`Coaching the 7th grade girls basketball team.`,
			`Dad coming home drunk.`,
			`Dad's big sex night.`,
			`Dad's famous chili.`,
			`Divorce.`,
			`Dolby® Digital Surround Sound.`,
			`Emotional unavailability.`,
			`Finding a place to sit down.`,
			`Flirting with the ladies at the bank.`,
			`Forty-two years of repressed homosexuality.`,
			`Going bald.`,
			`Having a bunch of kids by accident.`,
			`Having a heart attack.`,
			`Kidnapping Liam Neeson's daughter.`,
			`Making the printer work.`,
			`Puns.`,
			`Sitting on the toilet for 45 minutes.`,
			`Standing in the middle of the living room holding two remote controls.`,
			`Stealing a child's nose and keeping it forever.`,
			`The nipples of a man.`,
			`The son of a bitch who knocked up my daughter.`,
			`Tripping on an object and becoming angry.`,
			`What Dad has to say about Muslims.`,
		},
	}

	AddPack(pack)
}
