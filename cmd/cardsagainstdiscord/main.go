package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/cardsagainstdiscord"
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

	// Add slash command handler
	session.AddHandler(HandleInteractionCreate)
	
	// Add reaction and message handlers for game interaction
	session.AddHandler(func(s *discordgo.Session, ra *discordgo.MessageReactionAdd) {
		go cahManager.HandleReactionAdd(ra)
	})

	session.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
		go cahManager.HandleMessageCreate(msg)
	})

	err = session.Open()
	panicErr(err, "Failed opening gateway connection")
	log.Println("Connected to Discord!")

	// Register slash commands (this might take up to an hour to propagate globally)
	err = RegisterSlashCommands(session)
	if err != nil {
		log.Printf("Warning: Failed to register slash commands: %v. Slash commands may not work.", err)
	} else {
		log.Println("Slash commands registered successfully! Use /cah to interact with the bot.")
	}

	// We import http/pprof above to be able to inspect and do profiling
	go http.ListenAndServe(":7447", nil)
	select {}
}
