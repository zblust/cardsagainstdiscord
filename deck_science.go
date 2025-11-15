package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Science Pack",
		Description: "Science Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `A study published in Nature this week found that _ is good for you in small doses.`},
			&PromptCard{Prompt: `Hey there, Young Scientists! Put on your labcoats and strap on your safety goggles, because today we're learning about _!`},
			&PromptCard{Prompt: `In an attempt to recreate conditions just after the Big Bang, physicists at the LHC are observing collisions between _ and _.`},
			&PromptCard{Prompt: `In line with our predictions, we find a robust correlation between _ and _ (p<.05).`},
			&PromptCard{Prompt: `In what's being hailed as a major breakthrough, scientists have synthesized _ in the lab.`},
			&PromptCard{Prompt: `Today on Mythbusters, we found out how long _ can withstand _.`},
			&PromptCard{Prompt: `What really killed the dinosaurs?`},
		},

		Responses: []ResponseCard{
			`3.7 billion years of evolution.`,
			`A 0.7 waist-to-hip ratio.`,
			`A supermassive black hole.`,
			`Achieving reproductive success.`,
			`Being knowledgeable in a narrow domain that nobody understands or cares about.`,
			`David Attenborough watching us mate.`,
			`Developing secondary sex characteristics.`,
			`Driving into a tornado to learn about tornadoes.`,
			`Electroejaculating a capuchin monkey.`,
			`Evolving a labyrinthine vagina.`,
			`Explosive decompression.`,
			`Failing the Turing test.`,
			`Fun and interesting facts about rocks.`,
			`Getting really worried about global warming for a few seconds.`,
			`Infinity.`,
			`Insufficient serotonin.`,
			`Oxytocin release via manual stimulation of the nipples.`,
			`Photosynthesis.`,
			`Reconciling quantum theory with general relativity.`,
			`Slowly evaporating.`,
			`The quiet majesty of the sea turtle.`,
			`The Sun engulfing the Earth.`,
			`Urban legends`,
		},
	}

	AddPack(pack)
}
