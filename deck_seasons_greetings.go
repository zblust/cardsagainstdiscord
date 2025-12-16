package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "seasons-greetings-pack",
		Description: "Seasons Greetings Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Blood is thicker than %s.`},
			&PromptCard{Prompt: `Donna, pick up my dry cleaning and get my wife something for Christmas. I think she likes %s.`},
			&PromptCard{Prompt: `Here's what you can expect for the new year.
Out: %s.
In: %s.`},
			&PromptCard{Prompt: `It's beginning to look a lot like %s.`},
			&PromptCard{Prompt: `Jesus performed the miracle of %s but it was never recorded.`},
			&PromptCard{Prompt: `This holiday season, Tim Allen must overcome his fear of %s to save Christmas.`},
			&PromptCard{Prompt: `What's the one thing that makes an elf instantly ejaculate?`},
		},

		Responses: []ResponseCard{
			`A choir of angels descending from the sky and jizzing all over dad's sweater.`,
			`A frozen homeless man shattering on your doorstep.`,
			`A Toyota Corolla with a "Bring Our Jobs Back" Trumper sticker.`,
			`Another shot of morphine.`,
			`Eliminate liberal bias!`,
			`Finding out that Santa isn't real.`,
			`Fucking up Silent Night in front of 3,000 parents.`,
			`Gift wrapping a live hamster.`,
			`How cool it is that I love jesus and he loves me back.`,
			`How great of a blowjob Jesus could give.`,
			`How many drinks and Deborah has had.`,
			`Is snowman that contains the soul of my dead father.`,
			`My hot neighbor`,
			`My uncle who voted for Trump.`,
			`Piece of shit Christmas card with no money in them.`,
			`Pretending to be one of the guys but actually being the spider god.`,
			`Probably grandma's last Christmas, kids.`,
			`Snow falling gently on the frozen body of an orphan boy.`,
			`Sodomising the corpse of Ben Franklin.`,
			`Starting to see where ISIS is coming from.`,
			`The 9,000 children who starred to death today.`,
			`The shocking stupidity of the American public.`,
			`These new jeans that look so cool but are actually so bad.`,
		},
	}

	AddPack(pack)
}
