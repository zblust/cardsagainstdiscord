package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Theatre Pack",
		Description: "Theatre Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Alright everybody, HOLD!
Kelly, why is there _ on my stage?`},
			&PromptCard{Prompt: `Comedy = Tragedy + _.`},
			&PromptCard{Prompt: `Let's take it from the top, and remember, you are _. Show me  _(SAME CARD AGAIN)_.`},
			&PromptCard{Prompt: `Match-maker,
match-maker,
make me a match.
Find me _.`},
			&PromptCard{Prompt: `This season at Manhattan Theatre Club: "Who's afraid of _?"`},
		},

		Responses: []ResponseCard{
			`A dead salesman.`,
			`A Drama Desk Award for Outstanding Sound Design in a Play.`,
			`A fog machine.`,
			`A play about people I don't like doing things I don't care about.`,
			`Absolutely butchering Sondheim.`,
			`All 59 inches of Kristin Chenoweth.`,
			`An autographed headshot of Nathan Lane.`,
			`Being crushed to death by a stage light.`,
			`Brief male nudity.`,
			`Five miso soups, four seaweed salads, three soy burger dinners, two tofu dog platters, and one pasta with meatless meatballs.`,
			`Forgetting your lines, shitting your pants, and your pants falling down.`,
			`Improv comedy.`,
			`Killing Dad and fucking Mom.`,
			`Linda, 18 but wise beyond her years, achingly beautiful.`,
			`My whole family watching.`,
			`Narcissistic Personality Disorder.`,
			`Problematic depictions of Asian characters.`,
			`Rampant misogyny and sexual harassment.`,
			`Taking a year off to study Japanese puppet theatre.`,
			`The Phantom of the Opera.`,
			`The wickedly talented, one and only, Adele Dazeem.`,
			`This old lady next to me who won't stop farting.`,
			`Two contrasting monologues-- one classical, one contemporary.`,
			`Two men in a horse costume.`,
			`Two tickets to Hamilton.`,
		},
	}

	AddPack(pack)
}
