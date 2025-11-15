package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "PAX 2010 \"Oops\" Kit",
		Description: "PAX 2010 \"Oops\" Kit pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `_: Has science gone too far?`},
			&PromptCard{Prompt: `Daddy, why is Mommy crying?`},
			&PromptCard{Prompt: `Dear Agony Aunt, I'm having some trouble with _ and I need your advice.`},
			&PromptCard{Prompt: `I dropped my cellphone in the _ and now it won't work anymore.`},
			&PromptCard{Prompt: `Who ate my _ again?`},
		},
		Responses: []ResponseCard{
			`A Bitch Slap.`,
			`Chunks of food in my teeth`,
			`Extremely tight trousers.`,
			`One unforgettable night of passion.`,
			`Strafing an Afghan shepherd and his sad, sickly goats.`,
			`The boy who sucks the farts our of my sweatpants.`,
			`The entire Mormon Tabernacle Chior.`,
			`The few shreds of tinsel still clinging to my asshole.`,
			`Throwing another pidgey in the wood chipper.`,
		},
	}
	AddPack(pack)
}
