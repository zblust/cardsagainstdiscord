cards against discord

A discord bot for cards against humanity, and unlike other cah bots, you dont type your shit, you use the power of reactions to pick your shizz, MAGIC.

Pretty functional, some bugs may be around.

## Features

- **Modern Slash Commands**: Uses Discord's native slash commands (`/cah`) for easy discoverability
- **Single Dependency**: Built with the official bwmarrin/discordgo library
- **Reaction-Based Gameplay**: Use emoji reactions to join, play, and interact with the game

## Using the Bot

### Slash Commands

The bot uses Discord's native slash commands:

- `/cah help` - Shows help for all commands
- `/cah create [packs] [vote-mode]` - Creates a new game
- `/cah stop` - Stops the current game (game master only)
- `/cah kick <user>` - Kicks a player (game master only)
- `/cah packs` - Lists all available card packs

## Installation

To compile and run the standalone bot: (assuming you have go and git installed)

```
go get github.com/jonas747/cardsagainstdiscord/cmd/cardsagainstdiscord
cd $GOPATH/bin
# The executable now lies here, set DG_TOKEN to your token prefixed with "Bot " and run it
```

Or build from source:

```bash
git clone https://github.com/jonas747/cardsagainstdiscord
cd cardsagainstdiscord
go build ./cmd/cardsagainstdiscord
DG_TOKEN="Bot YOUR_TOKEN_HERE" ./cardsagainstdiscord
```

**Note**: When the bot starts, it will register slash commands with Discord. Global slash commands can take up to an hour to propagate, but they work immediately in servers where the bot has been manually added.

## Creating Custom Card Packs

Card packs are defined in `deck_*.go` files. Each prompt uses `%s` placeholders for cards that players need to pick.

### Same-Card References

If you want the same card to appear multiple times in a prompt, use position references (`%0`, `%1`, etc.) instead of multiple `%s` placeholders:

```go
// Old way (requires 3 picks for the same card):
&PromptCard{Prompt: `You want %s? You can't handle %s%s!`}

// New way (requires only 1 pick):
&PromptCard{Prompt: `You want %s? You can't handle %0!`}
```

- `%0` references the first card, `%1` references the second, etc.
- Players see `[FIRST CARD AGAIN]`, `[SECOND CARD AGAIN]`, etc. in the prompt
- Position references don't count toward the number of picks needed
- Supports up to 10 position references (%0 through %9)

### Example Pack

```go
package cardsagainstdiscord

func init() {
    pack := &CardPack{
        Name:        "mypack",
        Description: "My Custom Pack",
        Prompts: []*PromptCard{
            &PromptCard{Prompt: `I love %s.`},
            &PromptCard{Prompt: `You want %s? You can't handle %0!`},
        },
        Responses: []ResponseCard{
            `Puppies`,
            `Pizza`,
        },
    }
    AddPack(pack)
}
```
