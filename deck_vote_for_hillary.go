package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "vote-for-hillary",
		Description: "Vote For Hillary Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `As reparations for slavery, all African Americans will receive %s.`},
			&PromptCard{Prompt: `Senator, I trust you enjoyed %s last night.  Now, can I count on your vote?`},
			&PromptCard{Prompt: `When you go to the polls on Tuesday, remember:  a vote for me is a vote for %s.`},
		},

		Responses: []ResponseCard{
			`A beautiful, ever-expanding circle of inclusivity that will never include Republicans.`,
			`Black lives mattering.`,
			`Donald Trump holding his nose while he eats pussy.`,
			`Eating the president's pussy.`,
			`Increasing economic and political polarization.`,
			`Keeping the government out of my vagina.`,
			`Kicking the middle class in the balls with a regressive tax code.`,
			`Letting Bernie Sanders rest his world-weary head on your lap.`,
			`Slapping Ted Cruz over and over.`,
			`The Bernie Sanders revolution.`,
			`The fact that Hillary Clinton is a woman.`,
			`The systematic disenfranchisement of black voters.`,
		},
	}

	AddPack(pack)
}
