package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "cah-hidden-gems-bundle-a-few-new-cards-we-crammed-into-this-bundle-pack-amazon-exclusive",
		Description: "CAH: Hidden Gems Bundle: A Few New Cards We Crammed Into This Bundle Pack (Amazon Exclusive)",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `HELP WANTED: Need assistance with %s. No experience necessary.`},
			&PromptCard{Prompt: `Oh, whoops! These are not the right pants for %s.`},
			&PromptCard{Prompt: `You want %s? You can't handle %s[SAME CARD AGAIN]%s!`},
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
