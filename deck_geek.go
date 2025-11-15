package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Geek Pack",
		Description: "Geek Pack pack",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `_ is way better in _ mode.`},
			&PromptCard{Prompt: `_: Achievement unlocked.`},
			&PromptCard{Prompt: `(Heavy breathing) Luke, I am _.`},
			&PromptCard{Prompt: `Press ↓ ↓ ← → B to unleash _.`},
			&PromptCard{Prompt: `What made Spock cry?`},
			&PromptCard{Prompt: `What's the latest bullshit that's troubling this quaint fantasy town?`},
		},

		Responses: []ResponseCard{
			`A fully-dressed female videogame character.`,
			`A grumpy old Harrison Ford who'd rather be doing anything else.`,
			`A homemade, cum-stained Star Trek Uniform.`,
			`Achieving 500 actions per minute.`,
			`Charging up all the way.`,
			`Eating a pizza that's lying in the street to gain health.`,
			`Endless ninjas.`,
			`Forgetting to eat, and consequently dying.`,
			`Getting bitten by a radioactive spider and then battling leukemia for 30 years.`,
			`Getting bitten by a radioactive spider and then battling leukimia for 30 years.`,
			`KHAAAAAAAAN!`,
			`Loading from a previous save.`,
			`Offering sexual favors for an ore and a sheep.`,
			`Running out of stamina.`,
			`Separate drinking fountains for dark elves.`,
			`Ser Jorah Mormont's cerulean-blue balls.`,
			`Sharpening a foam broadsword on a foam whetstone.`,
			`Stuffing my balls into a Sega Genesis and pressing the power button.`,
			`Taking 2d6 emotional damage.`,
			`Tapping Serra Angel.`,
			`The Cock Ring of Alacrity.`,
			`The collective wail of every Magic player suddenly realizing that they've spent hundreds of dollars on pieces of cardboard.`,
			`The depression that ensues after catching 'em all.`,
			`Yoshi's huge egg-laying cloaca.`,
		},
	}

	AddPack(pack)
}
