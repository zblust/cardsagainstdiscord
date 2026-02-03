# ID Migration Example: int64 → string

This document shows concrete examples of how migrating from `int64` to `string` IDs would look when switching from the old discordgo fork to the official bwmarrin/discordgo library.

## Why the Change?

- **Old (jonas747/discordgo)**: Uses `int64` for Discord IDs (channels, guilds, users, messages)
- **New (bwmarrin/discordgo)**: Uses `string` for Discord IDs (Discord snowflakes)

## Example 1: Game Struct

### BEFORE (Current - int64):
```go
type Game struct {
	sync.RWMutex `json:"-" msgpack:"-"`
	Manager *GameManager       `json:"-" msgpack:"-"`
	Session *discordgo.Session `json:"-" msgpack:"-"`

	// The main channel this game resides in
	MasterChannel int64
	// The server the game resides in
	GuildID int64
	// The user that created this game
	GameMaster int64
	// The current cardzar
	CurrentCardCzar int64
	
	LastMenuMessage int64
	
	Players []*Player
	// ... other fields
}
```

### AFTER (Migrated - string):
```go
type Game struct {
	sync.RWMutex `json:"-" msgpack:"-"`
	Manager *GameManager       `json:"-" msgpack:"-"`
	Session *discordgo.Session `json:"-" msgpack:"-"`

	// The main channel this game resides in
	MasterChannel string
	// The server the game resides in
	GuildID string
	// The user that created this game
	GameMaster string
	// The current cardzar
	CurrentCardCzar string
	
	LastMenuMessage string
	
	Players []*Player
	// ... other fields
}
```

**Impact**: Every reference to these fields throughout the codebase needs updating.

---

## Example 2: GameManager

### BEFORE (Current - int64):
```go
type GameManager struct {
	sync.RWMutex
	SessionProvider SessionProvider
	ActiveGames     map[int64]*Game  // Maps channel/user IDs to games
	NumActiveGames  int
}

func (gm *GameManager) CreateGame(
	guildID int64,
	channelID int64,
	userID int64,
	username string,
	voteMode bool,
	packs ...string,
) (*Game, error) {
	// Check if game exists in channel
	if _, ok := gm.ActiveGames[channelID]; ok {
		return nil, ErrGameAlreadyInChannel
	}
	
	// Check if user is already in a game
	if _, ok := gm.ActiveGames[userID]; ok {
		return nil, ErrPlayerAlreadyInGame
	}
	
	game := &Game{
		MasterChannel: channelID,
		GuildID:       guildID,
		GameMaster:    userID,
		// ...
	}
	
	return game, nil
}
```

### AFTER (Migrated - string):
```go
type GameManager struct {
	sync.RWMutex
	SessionProvider SessionProvider
	ActiveGames     map[string]*Game  // Maps channel/user IDs to games
	NumActiveGames  int
}

func (gm *GameManager) CreateGame(
	guildID string,
	channelID string,
	userID string,
	username string,
	voteMode bool,
	packs ...string,
) (*Game, error) {
	// Check if game exists in channel
	if _, ok := gm.ActiveGames[channelID]; ok {
		return nil, ErrGameAlreadyInChannel
	}
	
	// Check if user is already in a game
	if _, ok := gm.ActiveGames[userID]; ok {
		return nil, ErrPlayerAlreadyInGame
	}
	
	game := &Game{
		MasterChannel: channelID,
		GuildID:       guildID,
		GameMaster:    userID,
		// ...
	}
	
	return game, nil
}
```

**Impact**: All function signatures accepting IDs need to change from `int64` to `string`.

---

## Example 3: Player Struct

### BEFORE (Current - int64):
```go
type Player struct {
	ID       int64
	Username string
	Channel  int64  // Private message channel
	InGame   bool
	// ... other fields
}

func (g *Game) AddPlayer(id int64, username string) bool {
	// Create DM channel
	channel, err := g.Session.UserChannelCreate(id)
	if err != nil {
		return false
	}
	
	player := &Player{
		ID:       id,
		Username: username,
		Channel:  channel.ID,
		InGame:   true,
		// ...
	}
	// ...
}
```

### AFTER (Migrated - string):
```go
type Player struct {
	ID       string
	Username string
	Channel  string  // Private message channel
	InGame   bool
	// ... other fields
}

func (g *Game) AddPlayer(id string, username string) bool {
	// Create DM channel
	channel, err := g.Session.UserChannelCreate(id)
	if err != nil {
		return false
	}
	
	player := &Player{
		ID:       id,
		Username: username,
		Channel:  channel.ID,
		InGame:   true,
		// ...
	}
	// ...
}
```

---

## Example 4: Message Handling

### BEFORE (Current - int64):
```go
func (g *Game) HandleReactionAdd(ra *discordgo.MessageReactionAdd) {
	channelID := ra.ChannelID  // Already int64
	userID := ra.UserID        // Already int64
	messageID := ra.MessageID  // Already int64
	
	if messageID != g.LastMenuMessage {
		return
	}
	
	player := g.findPlayer(userID)
	// ...
}

func (g *Game) findPlayer(id int64) *Player {
	for _, p := range g.Players {
		if p.ID == id {
			return p
		}
	}
	return nil
}
```

### AFTER (Migrated - string):
```go
func (g *Game) HandleReactionAdd(ra *discordgo.MessageReactionAdd) {
	channelID := ra.ChannelID  // Now string
	userID := ra.UserID        // Now string
	messageID := ra.MessageID  // Now string
	
	if messageID != g.LastMenuMessage {
		return
	}
	
	player := g.findPlayer(userID)
	// ...
}

func (g *Game) findPlayer(id string) *Player {
	for _, p := range g.Players {
		if p.ID == id {
			return p
		}
	}
	return nil
}
```

---

## Example 5: ID Comparisons

### BEFORE (Current - int64):
```go
func NextCardCzar(players []*Player, current int64) int64 {
	var next int64 = 0
	var lowest int64 = 0
	
	for _, v := range players {
		if !v.PlayingThisRound() {
			continue
		}
		
		if v.ID == current {
			continue
		}
		
		if v.ID > current && (v.ID < next || next == 0) {
			next = v.ID
		}
		
		if lowest == 0 || v.ID < lowest {
			lowest = v.ID
		}
	}
	
	if next == 0 {
		next = lowest
	}
	
	return next
}
```

### AFTER (Migrated - string):
```go
func NextCardCzar(players []*Player, current string) string {
	var next string = ""
	var lowest string = ""
	
	for _, v := range players {
		if !v.PlayingThisRound() {
			continue
		}
		
		if v.ID == current {
			continue
		}
		
		// String comparison for snowflakes works since they're chronological
		if v.ID > current && (v.ID < next || next == "") {
			next = v.ID
		}
		
		if lowest == "" || v.ID < lowest {
			lowest = v.ID
		}
	}
	
	if next == "" {
		next = lowest
	}
	
	return next
}
```

**Note**: String comparison of Discord snowflakes works because they're designed to be lexicographically sortable.

---

## Example 6: Session Methods

### BEFORE (Current - int64):
```go
// Send message to channel
msg, err := g.Session.ChannelMessageSendEmbed(g.MasterChannel, embed)
// g.MasterChannel is int64

// Create DM channel
channel, err := g.Session.UserChannelCreate(userID)
// userID is int64, channel.ID is int64

// Add reaction
err := g.Session.MessageReactionAdd(channelID, messageID, emoji)
// Both channelID and messageID are int64
```

### AFTER (Migrated - string):
```go
// Send message to channel
msg, err := g.Session.ChannelMessageSendEmbed(g.MasterChannel, embed)
// g.MasterChannel is string

// Create DM channel
channel, err := g.Session.UserChannelCreate(userID)
// userID is string, channel.ID is string

// Add reaction
err := g.Session.MessageReactionAdd(channelID, messageID, emoji)
// Both channelID and messageID are string
```

**Impact**: All Discord API method calls remain the same, but the types change.

---

## Files Requiring Changes

Based on the codebase structure, here are the main files that would need updates:

1. **game.go** (~100+ changes)
   - Game struct fields
   - All methods accepting/returning IDs
   - Message handling
   - Player management

2. **manager.go** (~50+ changes)
   - GameManager struct
   - CreateGame signature
   - All ID-based lookups
   - Active games map

3. **cad.go** (~20+ changes)
   - Helper functions
   - Error handling

4. **cmd/cardsagainstdiscord/main.go** (~30+ changes)
   - Command handlers
   - Remove old session code
   - Update all CreateGame calls

5. **cmd/cardsagainstdiscord/slashcommands.go**
   - Remove stringToInt64 conversion helper (no longer needed!)
   - Direct string passing

---

## Benefits After Migration

1. **Single Dependency**: Only bwmarrin/discordgo (actively maintained)
2. **No Type Conversions**: Slash commands already use strings
3. **Modern API**: Latest Discord features
4. **Cleaner Code**: No dual-session complexity

## Effort Estimate

- **Low-risk changes**: ~200 locations need `int64` → `string`
- **Search and replace**: Most changes are mechanical
- **Testing needed**: Game creation, player management, reactions
- **Time estimate**: 1-2 hours for experienced Go developer

---

## Current Dual-Session Approach (No Migration)

For comparison, here's what currently happens with dual sessions:

```go
// Slash command handler (bwmarrin/discordgo - string IDs)
func handleCreateCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channelID := i.ChannelID      // string
	userID := i.Member.User.ID    // string
	
	// Convert to int64 for game manager
	channelID64 := stringToInt64(channelID)
	userID64 := stringToInt64(userID)
	
	// Call game manager with int64
	cahManager.CreateGame(guildID64, channelID64, userID64, username, voteMode, packs...)
}

// Helper function needed for conversion
func stringToInt64(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}
```

With full migration, this conversion layer disappears entirely.
