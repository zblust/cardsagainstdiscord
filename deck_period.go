package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Period Pack",
		Description: "Period Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Can a woman really have it all? A career and _?`},
			&PromptCard{Prompt: `My body, my voice! _, my choice!`},
			&PromptCard{Prompt: `My vagina's angry. My vagina's furious and needs to talk. It needs to talk about _.`},
			&PromptCard{Prompt: `New fom Mattel, it's _ Barbie!`},
			&PromptCard{Prompt: `Tampax: Don't let your period ruin _.`},
			&PromptCard{Prompt: `What gets me wet?`},
		},

		Responses: []ResponseCard{
			`A diverse group of female friends casually discussing the side effects of birth control.`,
			`A woman president.`,
			`Always© Infinity Extra Heavy Overnight Pads with Wings.`,
			`An emotionally draining friendship.`,
			`Carrying a fetus to term.`,
			`Catching a whiff of my vag.`,
			`Dancing carefree in white linen pants.`,
			`Destroying a pair of underwear.`,
			`Drinking Beyonce's DivaCup and becoming immortal.`,
			`Driving my daughter to her abortion.`,
			`Eating three sleeves of Chips Ahoy!`,
			`Feeling bloaty and crampy.`,
			`Feeling lots of feelings.`,
			`Full bush.`,
			`How bloody that dick's about to be.`,
			`Masturbating with a Sonicare.`,
			`Period globs.`,
			`Playing with my pussy while I watch TV.`,
			`Post-sex funk.`,
			`Pulling out a never-ending tampon.`,
			`Pussy lips of all shapes and sizes.`,
			`Ringing out a sopping wet maxi pad into Donald Trump's mouth.`,
			`The vagina hole.`,
			`Using a Smucker's Uncrustable™ as a maxi pad.`,
		},
	}

	AddPack(pack)
}
