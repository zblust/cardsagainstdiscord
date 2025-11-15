package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Desert Bus For Hope Pack",
		Description: "Desert Bus For Hope Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Desert Bus: _ for the children.`},
			&PromptCard{Prompt: `I bless _ down in Africa.`},
			&PromptCard{Prompt: `Ken "_" Steacy went back into his archives and found his original sketches for _.`},
			&PromptCard{Prompt: `The four shifts at Desert Bus: Dawn Guard, Alpha Flight, Night Watch and _.`},
			&PromptCard{Prompt: `YOU try explaining _ to the media!`},
		},

		Responses: []ResponseCard{
			`$10,000 worth of silica gel.`,
			`A bathroom only for poop.`,
			`A chair-shaped fart sponge.`,
			`A lethal dose of caffeine.`,
			`Accidentally broadcasting an NSFW video to 5,000 people.`,
			`An encyclopedic knowledge of Night Court.`,
			`Belting out the chorus of a popular song and mumbling through the rest.`,
			`Bill's mom.`,
			`Desert Bus.`,
			`Doing it for the children.`,
			`Having no idea what the fuck is going on.`,
			`Letting the internet feed you.`,
			`Playing one-handed.`,
			`Shipping a fish brick to Ohio.`,
			`The BONE ZONE!`,
			`The D E V I C E.`,
			`The Turner Lickability Scale.`,
			`Whale dong.`,
			`William Shatner watching you.`,
		},
	}

	AddPack(pack)
}
