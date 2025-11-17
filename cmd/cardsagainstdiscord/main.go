package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sort"
	"strings"

	"github.com/jonas747/cardsagainstdiscord"
	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dstate"
)

var cahManager *cardsagainstdiscord.GameManager

func panicErr(err error, msg string) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

func main() {
	session, err := discordgo.New(os.Getenv("DG_TOKEN"))
	panicErr(err, "Failed initializing discordgo")

	cahManager = cardsagainstdiscord.NewGameManager(&cardsagainstdiscord.StaticSessionProvider{
		Session: session,
	})

	state := dstate.NewState()
	state.TrackMembers = false
	state.TrackPresences = false
	session.StateEnabled = false

	cmdSys := dcmd.NewStandardSystem("!cah")
	cmdSys.State = state
	cmdSys.Root.AddCommand(HelpCommand, dcmd.NewTrigger("help", "h").SetDisableInDM(true))
	cmdSys.Root.AddCommand(CreateGameCommand, dcmd.NewTrigger("create", "c").SetDisableInDM(true))
	cmdSys.Root.AddCommand(StopCommand, dcmd.NewTrigger("stop", "end", "s").SetDisableInDM(true))
	cmdSys.Root.AddCommand(KickCommand, dcmd.NewTrigger("kick").SetDisableInDM(true))
	cmdSys.Root.AddCommand(PacksCommand, dcmd.NewTrigger("packs", "p").SetDisableInDM(true))

	session.AddHandler(state.HandleEvent)
	session.AddHandler(cmdSys.HandleMessageCreate)
	session.AddHandler(func(s *discordgo.Session, ra *discordgo.MessageReactionAdd) {
		go cahManager.HandleReactionAdd(ra)
	})

	session.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
		go cahManager.HandleMessageCreate(msg)
	})

	err = session.Open()
	panicErr(err, "Failed opening gateway connection")
	log.Println("Running...")

	// We import http/pprof above to be ale to inspect shizz and do profiling
	go http.ListenAndServe(":7447", nil)
	select {}
}

var HelpCommand = &dcmd.SimpleCmd{
	ShortDesc: "Shows help for all Cards Against Humanity commands",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		help := "**Cards Against Humanity Bot - Commands**\n\n" +
			"**Game Management:**\n" +
			"➕ `!cah create [packs] [-v]` (aliases: `c`)\n" +
			"   Creates a new game in the current channel\n" +
			"   • `packs` - Space-separated pack names or `*` for all packs (default: `main`)\n" +
			"   • `-v` - Enable vote mode (no card czar, everyone votes)\n" +
			"   Examples:\n" +
			"   • `!cah create` - Start with main pack\n" +
			"   • `!cah create main bluebox -v` - Start with main and bluebox in vote mode\n" +
			"   • `!cah c *` - Start with all packs\n\n" +
			"➖ `!cah stop` (aliases: `end`, `s`)\n" +
			"   Stops the game in the current channel (game master only)\n" +
			"   Example: `!cah stop`\n\n" +
			"👢 `!cah kick <user>`\n" +
			"   Kicks a player from the game (game master only)\n" +
			"   • `user` - Mention the user to kick\n" +
			"   Example: `!cah kick @username`\n\n" +
			"📦 `!cah packs` (aliases: `p`)\n" +
			"   Lists all available card packs\n" +
			"   Example: `!cah packs`\n\n" +
			"ℹ️ `!cah help` (aliases: `h`)\n" +
			"   Shows this help message\n" +
			"   Example: `!cah help`\n\n" +
			"**In-Game Actions:**\n" +
			"• React with ➕ to join a game\n" +
			"• React with ➖ to leave a game\n" +
			"• React with ⏯ to start/pause (game master only)\n" +
			"• React with 🔄 to discard and redraw cards (during card selection)\n\n" +
			"**Gameplay Features:**\n" +
			"• **Card Discard/Redraw**: During your turn, react with 🔄 to enter discard mode, select cards to discard (marked with ✓), then react 🔄 again to confirm\n" +
			"• **Solo Testing**: Use `-v` flag to enable solo play in vote mode for testing\n" +
			"• All game communication happens via DM after joining\n\n" +
			"**Tips:**\n" +
			"• Use vote mode (`-v`) for casual/faster games or solo testing\n" +
			"• You can include multiple packs by separating them with spaces\n" +
			"• Game master can pause/resume with ⏯ and kick players\n" +
			"• Players only receive DMs while actively in the game"

		return help, nil
	},
}

var CreateGameCommand = &dcmd.SimpleCmd{
	ShortDesc: "Creates a cards against humanity game in this channel",
	CmdArgDefs: []*dcmd.ArgDef{
		&dcmd.ArgDef{Name: "packs", Type: dcmd.String, Default: "main", Help: "Packs seperated by space, or * to include all of them"},
	},
	CmdSwitches: []*dcmd.ArgDef{
		{Switch: "v", Name: "Vote mode, no cardczar"},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		voteMode := data.Switch("v").Bool()
		pStr := data.Args[0].Str()
		packs := strings.Fields(pStr)

		_, err := cahManager.CreateGame(data.GS.ID, data.CS.ID, data.Msg.Author.ID, data.Msg.Author.Username, voteMode, packs...)
		if err == nil {
			log.Println("Created a new game in ", data.CS.ID)
			return "", nil
		}

		if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
			return cahErr, nil
		}

		return "Something went wrong", err

	},
}

var StopCommand = &dcmd.SimpleCmd{
	ShortDesc: "Stops a cards against humanity game in this channel",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		err := cahManager.TryAdminRemoveGame(data.Msg.Author.ID)
		if err != nil {
			if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
				return cahErr, nil
			}

			return "Something went wrong", err
		}

		return "Stopped the game", nil
	},
}

var KickCommand = &dcmd.SimpleCmd{
	ShortDesc:       "Kicks a player from the card against humanity game in this channel, only the game master can do this",
	RequiredArgDefs: 1,
	CmdArgDefs: []*dcmd.ArgDef{
		&dcmd.ArgDef{Name: "user", Type: dcmd.UserID},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		userID := data.Args[0].Int64()

		err := cahManager.AdminKickUser(data.Msg.Author.ID, userID)
		if err != nil {
			if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
				return cahErr, nil
			}

			return "Something went wrong", err
		}

		return "User removed", nil
	},
}

var PacksCommand = &dcmd.SimpleCmd{
	ShortDesc: "Lists available packs",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		// Convert map to slice and sort by name (case-insensitive)
		packs := make([]*cardsagainstdiscord.CardPack, 0, len(cardsagainstdiscord.Packs))
		for _, v := range cardsagainstdiscord.Packs {
			packs = append(packs, v)
		}
		sort.Slice(packs, func(i, j int) bool {
			return strings.ToLower(packs[i].Name) < strings.ToLower(packs[j].Name)
		})

		resp := "Available packs:\n\n"
		for _, v := range packs {
			resp += "`" + v.Name + "` - " + v.Description + "\n"
		}

		return resp, nil
	},
}
