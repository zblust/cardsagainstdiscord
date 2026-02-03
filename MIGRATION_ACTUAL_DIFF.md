# Actual Git Diff Example: What the Changes Look Like

This shows what a real `git diff` would look like when migrating from int64 to string IDs.

## Example Diff for manager.go

```diff
diff --git a/manager.go b/manager.go
index 1234567..abcdefg 100644
--- a/manager.go
+++ b/manager.go
@@ -72,7 +72,7 @@ func ProcessPacks(packs ...string) ([]string, error) {
 type GameManager struct {
 sync.RWMutex
 SessionProvider SessionProvider
-ActiveGames     map[int64]*Game
+ActiveGames     map[string]*Game
 NumActiveGames  int
 }
 
@@ -80,7 +80,7 @@ func NewGameManager(sessionProvider SessionProvider) *GameManager {
 return &GameManager{
-ActiveGames:     make(map[int64]*Game),
+ActiveGames:     make(map[string]*Game),
 Provider: sessionProvider,
 }
 }
 
-func (gm *GameManager) CreateGame(guildID int64, channelID int64, userID int64, username string, voteMode bool, packs ...string) (*Game, error) {
+func (gm *GameManager) CreateGame(guildID string, channelID string, userID string, username string, voteMode bool, packs ...string) (*Game, error) {
 // Process packs using the helper function
 processedPacks, err := ProcessPacks(packs...)
 if err != nil {
@@ -130,7 +130,7 @@ func (gm *GameManager) CreateGame(guildID int64, channelID int64, userID int64,
 return game, err
 }
 
-func (gm *GameManager) FindGameFromChannelOrUser(id int64) *Game {
+func (gm *GameManager) FindGameFromChannelOrUser(id string) *Game {
 gm.RLock()
 defer gm.RUnlock()
 
@@ -141,7 +141,7 @@ func (gm *GameManager) FindGameFromChannelOrUser(id int64) *Game {
 return nil
 }
 
-func (gm *GameManager) PlayerTryJoinGame(gameID, playerID int64, username string) error {
+func (gm *GameManager) PlayerTryJoinGame(gameID, playerID string, username string) error {
 gm.Lock()
 defer gm.Unlock()
 
@@ -160,7 +160,7 @@ func (gm *GameManager) PlayerTryJoinGame(gameID, playerID int64, username strin
 return ErrGameNotFound
 }
 
-func (gm *GameManager) PlayerTryLeaveGame(playerID int64) error {
+func (gm *GameManager) PlayerTryLeaveGame(playerID string) error {
 gm.Lock()
 defer gm.Unlock()
 
@@ -173,7 +173,7 @@ func (gm *GameManager) PlayerTryLeaveGame(playerID int64) error {
 return ErrGameNotFound
 }
 
-func (gm *GameManager) AdminKickUser(admin, playerID int64) error {
+func (gm *GameManager) AdminKickUser(admin, playerID string) error {
 gm.Lock()
 defer gm.Unlock()
 
@@ -197,7 +197,7 @@ func (gm *GameManager) AdminKickUser(admin, playerID int64) error {
 return nil
 }
 
-func (gm *GameManager) RemoveGame(gameID int64) error {
+func (gm *GameManager) RemoveGame(gameID string) error {
 gm.Lock()
 defer gm.Unlock()
 
@@ -224,7 +224,7 @@ func (gm *GameManager) RemoveGame(gameID int64) error {
 return nil
 }
 
-func (gm *GameManager) TryAdminRemoveGame(admin int64) error {
+func (gm *GameManager) TryAdminRemoveGame(admin string) error {
 gm.Lock()
 defer gm.Unlock()
```

---

## Example Diff for game.go (Game struct)

```diff
diff --git a/game.go b/game.go
index 7654321..fedcba9 100644
--- a/game.go
+++ b/game.go
@@ -77,13 +77,13 @@ type Game struct {
 Session *discordgo.Session `json:"-" msgpack:"-"`
 
 // The main channel this game resides in, never changes
-MasterChannel int64
+MasterChannel string
 // The server the game resides in, never changes
-GuildID int64
+GuildID string
 
 // The user that created this game
-GameMaster int64
+GameMaster string
 
 // The current cardzar
-CurrentCardCzar int64
+CurrentCardCzar string
 
 PlayerLimit        int
 WinLimit           int
@@ -100,7 +100,7 @@ type Game struct {
 
 CurrentPropmpt *PromptCard
 
-LastMenuMessage int64
+LastMenuMessage string
 
 Responses []*PickedResonse
 
@@ -118,9 +118,9 @@ type PickedResonse struct {
 }
 
 type Player struct {
-ID       int64
+ID       string
 Username string
-Channel  int64
+Channel  string
 
-MessageID int64
+MessageID string
 
 Cards             []ResponseCard
```

---

## Example Diff for cmd/cardsagainstdiscord/slashcommands.go

This file gets SIMPLER after migration:

```diff
diff --git a/cmd/cardsagainstdiscord/slashcommands.go b/cmd/cardsagainstdiscord/slashcommands.go
index abc1234..def5678 100644
--- a/cmd/cardsagainstdiscord/slashcommands.go
+++ b/cmd/cardsagainstdiscord/slashcommands.go
@@ -173,12 +173,9 @@ func handleCreateCommand(s *discordgo.Session, i *discordgo.InteractionCreate,
 
 packs := strings.Fields(packsStr)
 
-// Convert IDs from string to int64
-channelID := stringToInt64(i.ChannelID)
-guildID := stringToInt64(i.GuildID)
-userID := stringToInt64(i.Member.User.ID)
+// IDs are already strings - no conversion needed!
 username := i.Member.User.Username
 
-_, err := cahManager.CreateGame(guildID, channelID, userID, username, voteMode, packs...)
+_, err := cahManager.CreateGame(i.GuildID, i.ChannelID, i.Member.User.ID, username, voteMode, packs...)
 if err != nil {
 cahErr := cardsagainstdiscord.HumanizeError(err); cahErr != "" {
 d(s, i, cahErr)
@@ -193,7 +190,7 @@ func handleCreateCommand(s *discordgo.Session, i *discordgo.InteractionCreate,
 }
 
 func handleStopCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
-userID := stringToInt64(i.Member.User.ID)
+userID := i.Member.User.ID
 
 err := cahManager.TryAdminRemoveGame(userID)
 if err != nil {
@@ -215,8 +212,8 @@ func handleKickCommand(s *discordgo.Session, i *discordgo.InteractionCreate, op
 
 }
 
-targetUserID := stringToInt64(options[0].UserValue(s).ID)
-adminID := stringToInt64(i.Member.User.ID)
+targetUserID := options[0].UserValue(s).ID
+adminID := i.Member.User.ID
 
 err := cahManager.AdminKickUser(adminID, targetUserID)
 if err != nil {
@@ -271,10 +268,3 @@ func respondError(s *discordgo.Session, i *discordgo.InteractionCreate, message
 tf("Error responding to interaction: %v", err)
 }
 }
-
-// stringToInt64 converts a Discord snowflake ID string to int64
-func stringToInt64(s string) int64 {
-var id int64
-fmt.Sscanf(s, "%d", &id)
-return id
-}
```

**Note**: The slash commands file gets SHORTER and SIMPLER! No more conversion needed.

---

## Statistics from Example Diffs

From the examples above:

- **Lines changed**: ~200-250 across all files
- **Files affected**: 4-5 main files
- **Lines removed**: ~10 (conversion helper and its uses)
- **Complexity removed**: Type conversion layer eliminated

## How to Generate Full Diff

If you want to see all changes at once:

```bash
# On a migration branch:
git diff main -- '*.go'

# Count the changes:
git diff main --stat
```

Example output would be:
```
 cad.go                               |  15 ++--
 game.go                              | 123 +++++++++++++++--------------
 manager.go                           |  52 ++++++-------
 cmd/cardsagainstdiscord/main.go      |  35 ++++-----
 cmd/cardsagainstdiscord/slashcommands.go | 23 +-----
 5 files changed, 115 insertions(+), 133 deletions(-)
```

Notice: Fewer total lines! (133 deletions vs 115 insertions = -18 lines net)
