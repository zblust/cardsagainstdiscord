package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "2014 Holiday Pack",
		Description: "2014 Holiday Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `A curse upon thee! Many years from now, just when you think you're safe, _ shall turn into _.`},
			&PromptCard{Prompt: `Behold the Four Horsemen of the Apocalypse! War, Famine, Death, and _.`},
			&PromptCard{Prompt: `Dear Mom and Dad, Camp is fun. I like capture the flag. Yesterday, one of the older kids taught me about _. I love you, Casey`},
			&PromptCard{Prompt: `Here lies First Last, Year-2015, devoted friend, lover of _.`},
			&PromptCard{Prompt: `Honey, Mommy and Daddy love you very much. But apparently Mommy loves _ more than she loves Daddy.`},
			&PromptCard{Prompt: `Today on Buzzfeed: 10 pictures of _ That look like _!`},
			&PromptCard{Prompt: `Why am I so tired?`},
		},

		Responses: []ResponseCard{
			`200 years of slavery.`,
			`A cloud of ash that darkens the Earth for a thousand years.`,
			`A protracted siege.`,
			`A vague fear of something called ISIS.`,
			`All the poop inside of my body.`,
			`Being replaced by a robot.`,
			`Blockbuster late fees up the wazoo.`,
			`Building a ladder of hot dogs to the moon.`,
			`Ebola.`,
			`Harnessing the miraculous power of the atom to slaughter 200,000 Japanese people.`,
			`Reading an entire book.`,
			`Rising sea levels consistent with scientific predictions.`,
			`Rock music and premarital sex.`,
			`Small-town cops with M-4 assault rifles.`,
			`The 9,000 children who starved to death today.`,
			`The Bowflex Revolution.`,
			`The diminishing purity of the white race.`,
			`The dying breath of the last human.`,
			`The events depicted in James Cameron's Avatar.`,
			`The Great Lizard Uprising of 2352.`,
			`The transience of all things.`,
			`This groovy new thing called LSD.`,
			`Trying to feel something, anything.`,
			`What remains of my penis.`,
		},
	}

	AddPack(pack)
}
