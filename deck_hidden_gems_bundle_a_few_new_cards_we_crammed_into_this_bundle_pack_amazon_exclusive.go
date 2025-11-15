package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "CAH: Hidden Gems Bundle: A Few New Cards We Crammed Into This Bundle Pack (Amazon Exclusive)",
		Description: "CAH: Hidden Gems Bundle: A Few New Cards We Crammed Into This Bundle Pack (Amazon Exclusive) pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `HELP WANTED: Need assistance with _. No experience necessary.`},
			&PromptCard{Prompt: `Oh, whoops! These are not the right pants for _.`},
			&PromptCard{Prompt: `You want _? You can't handle _[SAME CARD AGAIN]_!`},
		},

		Responses: []ResponseCard{
			`A juicy lil' booty going poot-poot-pooty.`,
			`Aborted boat babies going overboard.`,
			`Dumping Mountain Dew Baja Blas all over my gorgeous natural titties.`,
			`Getting radicalized on YouTube.`,
			`Getting spit on by one hundred women.`,
			`Sitting atop a pile of tuna like some kind of tuna queen.`,
			`Watching the show Frasier" and feeling the emotion "pleasure."`,
		},
	}

	AddPack(pack)
}
