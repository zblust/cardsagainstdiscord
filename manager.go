package cardsagainstdiscord

import (
	"github.com/jonas747/discordgo"
	"sync"
)

// ProcessPacks processes pack arguments and returns the final list of pack names
// This is exposed for testing purposes
func ProcessPacks(packs ...string) ([]string, error) {
	allPacks := false
	whitelist := make([]string, 0)
	blacklist := make(map[string]bool)
	
	// First pass: identify all packs, whitelisted packs, and blacklisted packs
	for _, v := range packs {
		if v == "*" {
			allPacks = true
			continue
		}
		
		// Check if this is a blacklist entry (starts with -)
		if len(v) > 1 && v[0] == '-' {
			packName := v[1:] // Remove the - prefix
			_, ok := Packs[packName]
			if !ok {
				return nil, &ErrUnknownPack{
					PassedPack: v,
				}
			}
			blacklist[packName] = true
		} else {
			// Whitelist entry
			_, ok := Packs[v]
			if !ok {
				return nil, &ErrUnknownPack{
					PassedPack: v,
				}
			}
			whitelist = append(whitelist, v)
		}
	}

	// Determine final pack list
	var resultPacks []string
	if allPacks {
		// Start with all packs, then remove blacklisted ones
		resultPacks = make([]string, 0, len(Packs))
		for k := range Packs {
			if !blacklist[k] {
				resultPacks = append(resultPacks, k)
			}
		}
	} else if len(whitelist) > 0 {
		// Use whitelist (blacklist is ignored when not using *)
		resultPacks = whitelist
	} else if len(blacklist) > 0 {
		// Only blacklist provided without *, treat as error
		return nil, ErrNoPacks
	} else {
		// No packs specified
		return nil, ErrNoPacks
	}
	
	if len(resultPacks) < 1 {
		return nil, ErrNoPacks
	}
	
	return resultPacks, nil
}

type GameManager struct {
	sync.RWMutex
	SessionProvider SessionProvider
	ActiveGames     map[int64]*Game
	NumActiveGames  int
}

func NewGameManager(sessionProvider SessionProvider) *GameManager {
	return &GameManager{
		ActiveGames:     make(map[int64]*Game),
		SessionProvider: sessionProvider,
	}
}

func (gm *GameManager) CreateGame(guildID int64, channelID int64, userID int64, username string, voteMode bool, packs ...string) (*Game, error) {
	// Process packs using the helper function
	processedPacks, err := ProcessPacks(packs...)
	if err != nil {
		return nil, err
	}
	
	packs = processedPacks

	gm.Lock()
	defer gm.Unlock()

	if _, ok := gm.ActiveGames[channelID]; ok {
		return nil, ErrGameAlreadyInChannel
	}

	if _, ok := gm.ActiveGames[userID]; ok {
		return nil, ErrPlayerAlreadyInGame
	}

	game := &Game{
		MasterChannel: channelID,
		Manager:       gm,
		GuildID:       guildID,
		Packs:         packs,
		GameMaster:    userID,
		VoteMode:      voteMode,
		PlayerLimit:   20,
		WinLimit:      10,
		Session:       gm.SessionProvider.SessionForGuild(guildID),
	}

	err = game.Created()
	if err == nil {
		game.AddPlayer(userID, username)

		gm.ActiveGames[channelID] = game
		gm.ActiveGames[userID] = game
		gm.NumActiveGames++
	}

	return game, err
}

func (gm *GameManager) FindGameFromChannelOrUser(id int64) *Game {
	gm.RLock()
	defer gm.RUnlock()

	if g, ok := gm.ActiveGames[id]; ok {
		return g
	}

	return nil
}

func (gm *GameManager) PlayerTryJoinGame(gameID, playerID int64, username string) error {
	gm.Lock()
	defer gm.Unlock()

	if _, ok := gm.ActiveGames[playerID]; ok {
		return ErrPlayerAlreadyInGame
	}

	if g, ok := gm.ActiveGames[gameID]; ok {
		if g.AddPlayer(playerID, username) {
			gm.ActiveGames[playerID] = g
			return nil
		}

		return ErrGameFull
	}

	return ErrGameNotFound
}

func (gm *GameManager) PlayerTryLeaveGame(playerID int64) error {
	gm.Lock()
	defer gm.Unlock()

	if g, ok := gm.ActiveGames[playerID]; ok {
		delete(gm.ActiveGames, playerID)
		g.RemovePlayer(playerID)
		return nil
	}

	return ErrGameNotFound
}

func (gm *GameManager) AdminKickUser(admin, playerID int64) error {
	gm.Lock()
	defer gm.Unlock()

	g, ok := gm.ActiveGames[admin]
	if !ok {
		return ErrGameNotFound
	}

	g.RLock()
	if g.GameMaster != admin {
		g.RUnlock()
		return ErrNotGM
	}
	g.RUnlock()

	if g.RemovePlayer(playerID) {
		delete(gm.ActiveGames, playerID)
	} else {
		return ErrPlayerNotInGame
	}

	return nil
}

func (gm *GameManager) RemoveGame(gameID int64) error {
	gm.Lock()
	defer gm.Unlock()

	g, ok := gm.ActiveGames[gameID]
	if !ok {
		return ErrGameNotFound
	}

	g.Stop()

	// Remove all references to the game
	g.RLock()
	defer g.RUnlock()

	delete(gm.ActiveGames, g.MasterChannel)
	delete(gm.ActiveGames, g.GameMaster)
	for _, v := range g.Players {
		if v.InGame {
			delete(gm.ActiveGames, v.ID)
		}
	}

	gm.NumActiveGames--

	return nil
}

func (gm *GameManager) TryAdminRemoveGame(admin int64) error {
	gm.Lock()
	defer gm.Unlock()

	g, ok := gm.ActiveGames[admin]
	if !ok {
		return ErrGameNotFound
	}

	g.Lock()
	defer g.Unlock()

	if g.GameMaster != admin {
		return ErrNotGM
	}

	if g.stopped {
		return ErrStoppedAlready
	}

	close(g.stopch)
	g.stopped = true

	// Remove all references to the game
	delete(gm.ActiveGames, g.MasterChannel)
	delete(gm.ActiveGames, g.GameMaster)
	for _, v := range g.Players {
		if v.InGame {
			delete(gm.ActiveGames, v.ID)
		}
	}

	gm.NumActiveGames--

	return nil
}

func (gm *GameManager) HandleReactionAdd(ra *discordgo.MessageReactionAdd) {
	cid := ra.ChannelID
	userID := ra.UserID

	gm.RLock()
	if game, ok := gm.ActiveGames[cid]; ok {
		gm.RUnlock()
		game.HandleRectionAdd(ra)
	} else if game, ok := gm.ActiveGames[userID]; ok {
		gm.RUnlock()
		game.HandleRectionAdd(ra)
	} else {
		gm.RUnlock()
	}
}

func (gm *GameManager) HandleMessageCreate(msgCreate *discordgo.MessageCreate) {
	userID := msgCreate.Author.ID

	gm.RLock()
	if game, ok := gm.ActiveGames[userID]; ok {
		gm.RUnlock()
		game.HandleMessageCreate(msgCreate)
	} else {
		gm.RUnlock()
	}
}

func (gm *GameManager) LoadGameFromSerializedState(game *Game) {
	game.Session = gm.SessionProvider.SessionForGuild(game.GuildID)
	game.Manager = gm
	game.stopch = make(chan bool)

	gm.Lock()
	for _, v := range game.Players {
		if v.InGame {
			gm.ActiveGames[v.ID] = game
		}
	}

	gm.ActiveGames[game.MasterChannel] = game
	gm.NumActiveGames++
	gm.Unlock()

	game.loadFromSerializedState()
}
