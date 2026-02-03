package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/cardsagainstdiscord"
)

// Slash commands definitions
var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "cah",
			Description: "Cards Against Humanity bot commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "help",
					Description: "Shows help for all Cards Against Humanity commands",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "create",
					Description: "Creates a new game in the current channel",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "packs",
							Description: "Packs separated by space, or * for all (default: main). Use - prefix to exclude",
							Required:    false,
						},
						{
							Type:        discordgo.ApplicationCommandOptionBoolean,
							Name:        "vote-mode",
							Description: "Enable vote mode (no card czar, everyone votes)",
							Required:    false,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "stop",
					Description: "Stops the game in the current channel (game master only)",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "kick",
					Description: "Kicks a player from the game (game master only)",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user",
							Description: "The user to kick",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "packs",
					Description: "Lists all available card packs",
				},
			},
		},
	}
)

// RegisterSlashCommands registers all slash commands with Discord
func RegisterSlashCommands(s *discordgo.Session) error {
	// Note: We register commands globally. For development, you might want to use guild-specific commands
	// which update instantly instead of taking up to an hour
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			return fmt.Errorf("cannot create '%s' command: %v", cmd.Name, err)
		}
		log.Printf("Registered slash command: /%s", cmd.Name)
	}
	return nil
}

// HandleInteractionCreate handles slash command interactions
func HandleInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Extract command data
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	
	// Only handle /cah commands
	if data.Name != "cah" {
		return
	}

	// Get the subcommand
	if len(data.Options) == 0 {
		respondError(s, i, "No subcommand provided")
		return
	}

	subcommand := data.Options[0]

	// Route to appropriate handler
	switch subcommand.Name {
	case "help":
		handleHelpCommand(s, i)
	case "create":
		handleCreateCommand(s, i, subcommand.Options)
	case "stop":
		handleStopCommand(s, i)
	case "kick":
		handleKickCommand(s, i, subcommand.Options)
	case "packs":
		handlePacksCommand(s, i)
	default:
		respondError(s, i, "Unknown subcommand")
	}
}

func handleHelpCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	help := "**Cards Against Humanity Bot - Commands**\n\n" +
		"**Game Management:**\n" +
		"➕ `/cah create [packs] [vote-mode]`\n" +
		"   Creates a new game in the current channel\n" +
		"   • `packs` - Space-separated pack names or `*` for all packs (default: `main`)\n" +
		"   • Prefix pack names with `-` to exclude them (blacklist)\n" +
		"   • `vote-mode` - Enable vote mode (no card czar, everyone votes)\n" +
		"   Examples:\n" +
		"   • `/cah create` - Start with main pack\n" +
		"   • `/cah create packs:main bluebox vote-mode:true` - Start with main and bluebox in vote mode\n" +
		"   • `/cah create packs:*` - Start with all packs\n" +
		"   • `/cah create packs:* -weed -trump` - Start with all packs except weed and trump\n\n" +
		"➖ `/cah stop`\n" +
		"   Stops the game in the current channel (game master only)\n\n" +
		"👢 `/cah kick <user>`\n" +
		"   Kicks a player from the game (game master only)\n" +
		"   • `user` - Select the user to kick\n\n" +
		"📦 `/cah packs`\n" +
		"   Lists all available card packs\n\n" +
		"ℹ️ `/cah help`\n" +
		"   Shows this help message\n\n" +
		"**In-Game Actions:**\n" +
		"• React with ➕ to join a game\n" +
		"• React with ➖ to leave a game\n" +
		"• React with ⏯ to start/pause (game master only)\n" +
		"• React with 🔄 to discard and redraw cards (during card selection)\n\n" +
		"**Gameplay Features:**\n" +
		"• **Card Discard/Redraw**: During your turn, react with 🔄 to enter discard mode, select cards to discard (marked with ✓), then react 🔄 again to confirm\n" +
		"• **Solo Testing**: Use `vote-mode` to enable solo play for testing\n" +
		"• All game communication happens via DM after joining\n\n" +
		"**Tips:**\n" +
		"• Use vote mode for casual/faster games or solo testing\n" +
		"• You can include multiple packs by separating them with spaces\n" +
		"• Game master can pause/resume with ⏯ and kick players\n" +
		"• Players only receive DMs while actively in the game"

	respond(s, i, help)
}

func handleCreateCommand(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) {
	// Parse options
	packsStr := "main"
	voteMode := false

	for _, opt := range options {
		switch opt.Name {
		case "packs":
			packsStr = opt.StringValue()
		case "vote-mode":
			voteMode = opt.BoolValue()
		}
	}

	packs := strings.Fields(packsStr)

	// Convert IDs from string to int64
	channelID := stringToInt64(i.ChannelID)
	guildID := stringToInt64(i.GuildID)
	userID := stringToInt64(i.Member.User.ID)
	username := i.Member.User.Username

	_, err := cahManager.CreateGame(guildID, channelID, userID, username, voteMode, packs...)
	if err != nil {
		if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
			respond(s, i, cahErr)
			return
		}
		respondError(s, i, "Something went wrong")
		log.Printf("Error creating game: %v", err)
		return
	}

	log.Println("Created a new game in", i.ChannelID, "via slash command")
	respond(s, i, "Game created! React with ➕ to join.")
}

func handleStopCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := stringToInt64(i.Member.User.ID)

	err := cahManager.TryAdminRemoveGame(userID)
	if err != nil {
		if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
			respond(s, i, cahErr)
			return
		}
		respondError(s, i, "Something went wrong")
		log.Printf("Error stopping game: %v", err)
		return
	}

	respond(s, i, "Stopped the game")
}

func handleKickCommand(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) {
	if len(options) == 0 {
		respondError(s, i, "No user specified")
		return
	}

	targetUserID := stringToInt64(options[0].UserValue(s).ID)
	adminID := stringToInt64(i.Member.User.ID)

	err := cahManager.AdminKickUser(adminID, targetUserID)
	if err != nil {
		if cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
			respond(s, i, cahErr)
			return
		}
		respondError(s, i, "Something went wrong")
		log.Printf("Error kicking user: %v", err)
		return
	}

	respond(s, i, "User removed")
}

func handlePacksCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	respond(s, i, resp)
}

// Helper functions

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}

func respondError(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "❌ " + message,
			Flags:   discordgo.MessageFlagsEphemeral, // Only visible to the user
		},
	})
	if err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}

// stringToInt64 converts a Discord snowflake ID string to int64
func stringToInt64(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}
