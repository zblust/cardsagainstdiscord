package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "reject-pack",
		Description: "Reject Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `From WBEZ Chicago, it's This American Life. Today on our program, %s. Stay with us.`},
			&PromptCard{Prompt: `My name is Inigo Montoya. You killed my father. Prepare for %s.`},
			&PromptCard{Prompt: `Sir, we found you passed out naked on the side of the road. What's the last thing you remember?`},
			&PromptCard{Prompt: `The elders of the Ibo tribe of Nigeria recommend %s as a cure for impotence.`},
			&PromptCard{Prompt: `The Westboro Baptist Church is now picketing soldiers' funerals with signs that read 'GOD HATES %s!'`},
			&PromptCard{Prompt: `What are two cards in your hand that you want to get rid of?`},
			&PromptCard{Prompt: `What do you see? [An image of a Rorschach inkblot]`},
			&PromptCard{Prompt: `You can't wait forever. It's time to talk to your doctor about %s.`},
		},
		Responses: []ResponseCard{
			`A giant squid in a wedding gown.`,
			`A heart that is two sizes too small and that therefore cannot pump an adequate amount of blood.`,
			`A sexy naked interactive theater thing.`,
			`Asshole pomegranates that are hard to eat.`,
			`Becoming so rich that you shed your body and turn to vapor.`,
			`Carribbean Jesus.`,
			`Cornhole 101: just drop it in the hole.`,
			`Crayon-colored vomit`,
			`Dividing by zero.`,
			`Ejaculating a pound of tinsel.`,
			`Faking a Mental Disorder`,
			`Ladles.`,
			`My dickhead roomate.`,
			`Playing an ocarina to summon Ultra-Congress from the sea.`,
			`Super yoga.`,
			`The John D. and Catherine T. MacArthur Foundation.`,
		},
	}
	AddPack(pack)
}
