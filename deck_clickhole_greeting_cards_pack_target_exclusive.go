package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "ClickHole Greeting Cards Pack (Target Exclusive)",
		Description: "ClickHole Greeting Cards Pack (Target Exclusive) pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Hallmark has invented a new holiday. It's called "_ Day."`},
			&PromptCard{Prompt: `Happy Birthday! I love you. Here's _.`},
			&PromptCard{Prompt: `Honey, it's gonna take more than a shitty greeting card to make up for ten years of _.`},
			&PromptCard{Prompt: `I cannot stop smiling about _.`},
			&PromptCard{Prompt: `Roses are red. Violets are blue. I am _ and so are you.`},
		},

		Responses: []ResponseCard{
			`A one-star Uber driver named "Wife of Tarantula."`,
			`An explosion of glitter.`,
			`Eating my children.`,
			`Emerging from my mother's vagina.`,
			`Give me kiss.`,
			`Looking like shit.`,
			`Luring my husband to the bedroom with a trail of ferret teeth and appearing on the bed in a ferret costume.`,
			`Mom's special birthday rimjob.`,
			`My five dead husbands.`,
			`Turning 70 and still being fuckable.`,
		},
	}

	AddPack(pack)
}
