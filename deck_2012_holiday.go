package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "2012 Holiday Pack",
		Description: "2012 Holiday Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `After blacking out during New Year's Eve, I was awoken by _.`},
			&PromptCard{Prompt: `Every Christmas, my uncle gets drunk and tells the story about _.`},
			&PromptCard{Prompt: `Jesus is _.`},
			&PromptCard{Prompt: `On the third day of Christmas, my true love gave to me: three French hens, two turtle doves, and _.`},
			&PromptCard{Prompt: `This is a little embarrassing but I was wondering if you could recommend a doctor who specializes in _.`},
			&PromptCard{Prompt: `Wake up, America. Christmas is under attack by secular liberals and their _.`},
			&PromptCard{Prompt: `What keeps me warm during the cold, cold winter?`},
		},

		Responses: []ResponseCard{
			`A Christmas stocking full of coleslaw.`,
			`A Hungry-Man™ Frozen Christmas Dinner for One.`,
			`A toxic family environment.`,
			`A visually arresting turtleneck.`,
			`Another shitty year.`,
			`Clearing a bloody path through Walmart with a scimitar.`,
			`Eating an entire snowman.`,
			`Elf cum.`,
			`Fucking up "Silent Night" in front of 300 parents.`,
			`Gift-wrapping a live hamster.`,
			`Immaculate conception.`,
			`Krampus, the Austrian Christmas monster.`,
			`Mall Santa.`,
			`My hot cousin.`,
			`Pretending to be happy.`,
			`Santa's heavy sack.`,
			`Several intertwining love stories featuring Hugh Grant.`,
			`Socks.`,
			`Space Jam on VHS.`,
			`Taking down Santa with a surface-to-air missile.`,
			`The Star Wars Holiday Special.`,
			`The tiny, calloused hands of the Chinese children that made this card.`,
			`Whatever Kwanzaa is supposed to be about.`,
		},
	}

	AddPack(pack)
}
