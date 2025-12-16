package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "theatre-pack---cats-musical-pack",
		Description: "Theatre Pack - CATS Musical Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `%s! All alone in the moonlight!`},
			&PromptCard{Prompt: `Ahhh %s! Toodle pip!`},
			&PromptCard{Prompt: `Jellicle Cats come out tonight, Jellicle Cats come one, come all, the time for %s is now, Jellicles come to the Jellicle Ball.`},
		},

		Responses: []ResponseCard{
			`30 grown adults crawling around on all-fours.`,
			`A man who has not heard of a Jellicle cat.`,
			`Becoming Mr. Mistoffelees.`,
			`Bustopher Jones.`,
			`Carbucketty.`,
			`Feline AIDS.`,
			`Garfield.`,
			`Griddlebone.`,
			`Grizabella, the glamour cat.`,
			`Growltiger.`,
			`Grump Cat.`,
			`Gus: the Palsy Cat.`,
			`Heathcliff.`,
			`Jellicle cats.`,
			`Macavity!`,
			`Memory.`,
			`Meow.`,
			`Mungojerrie and Rumpleteazer.`,
			`Munkustrap.`,
			`Old Deuteronomy.`,
			`Peter, Augustus, Alonzo, or James.`,
			`Practical cats,
dramatical cats,
oratical cats,
sceptical cats,
romantical cats,
parasitical cats,
statistical cats,
political cats,
hypocritical cats,
clerical cats,
cynical cats,
rabbinical cats.`,
			`Rum Tum Tugger.`,
			`Shimbleshanks.`,
			`Sir Andrew Lloyd Webber.`,
			`The mystical divinity of unashamed felinity.`,
			`Winkles and shrimps.`,
		},
	}

	AddPack(pack)
}
