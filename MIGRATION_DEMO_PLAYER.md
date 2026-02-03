# Concrete Example: Migrating Player-Related Code

This shows the **actual changes** needed for player-related functionality when migrating from int64 to string IDs.

## File: game.go (Player Management Section)

### CURRENT CODE (Using int64):

```go
type Player struct {
ID       int64
Username string
Channel  int64

MessageID int64

Cards             []ResponseCard
SelectedCards     []int
SelectedCardsLock sync.Mutex

Wins int

DiscardMode bool

InGame  bool
Playing bool
}

func (g *Game) AddPlayer(id int64, username string) bool {
g.Lock()
defer g.Unlock()

if len(g.Players) >= g.PlayerLimit {
 false
}

for _, v := range g.Players {
v.ID == id {
v.InGame {
 false
Game = true
ing = true
 true
nel, err := g.Session.UserChannelCreate(id)
if err != nil {
tln("Failed creating user channel", err)
 false
}

player := &Player{
      id,
ame: username,
nel:  channel.ID,
Game:   true,
ing:  true,
}

player.Cards = []ResponseCard{}

g.Players = append(g.Players, player)

return true
}

func (g *Game) RemovePlayer(id int64) bool {
g.Lock()
defer g.Unlock()

player := g.findPlayer(id)
if player == nil {
 false
}

player.InGame = false
player.Playing = false

numPlaying := 0
for _, v := range g.Players {
v.PlayingThisRound() {
umPlaying++
numPlaying < 1 {
g.Manager.RemoveGame(g.MasterChannel)
 true
}

if g.GameMaster == id && numPlaying > 0 {
_, v := range g.Players {
v.PlayingThisRound() {
= v.ID
g.sendAnnouncment(fmt.Sprintf("GameMaster left, assigned <@%d> as new game master.", v.ID), false)
 true
}

func (g *Game) findPlayer(id int64) *Player {
for _, v := range g.Players {
v.ID == id {
 v
 nil
}
```

### MIGRATED CODE (Using string):

```go
type Player struct {
ID       string  // CHANGED: int64 → string
Username string
Channel  string  // CHANGED: int64 → string

MessageID string  // CHANGED: int64 → string

Cards             []ResponseCard
SelectedCards     []int
SelectedCardsLock sync.Mutex

Wins int

DiscardMode bool

InGame  bool
Playing bool
}

func (g *Game) AddPlayer(id string, username string) bool {  // CHANGED: int64 → string
g.Lock()
defer g.Unlock()

if len(g.Players) >= g.PlayerLimit {
 false
}

for _, v := range g.Players {
v.ID == id {  // String comparison (works the same)
v.InGame {
 false
Game = true
ing = true
 true
nel, err := g.Session.UserChannelCreate(id)  // Still accepts string in new API
if err != nil {
tln("Failed creating user channel", err)
 false
}

player := &Player{
      id,        // Now string
ame: username,
nel:  channel.ID,  // channel.ID is now string
Game:   true,
ing:  true,
}

player.Cards = []ResponseCard{}

g.Players = append(g.Players, player)

return true
}

func (g *Game) RemovePlayer(id string) bool {  // CHANGED: int64 → string
g.Lock()
defer g.Unlock()

player := g.findPlayer(id)  // Call with string
if player == nil {
 false
}

player.InGame = false
player.Playing = false

numPlaying := 0
for _, v := range g.Players {
v.PlayingThisRound() {
umPlaying++
numPlaying < 1 {
g.Manager.RemoveGame(g.MasterChannel)  // MasterChannel now string
 true
}

if g.GameMaster == id && numPlaying > 0 {  // String comparison
_, v := range g.Players {
v.PlayingThisRound() {
= v.ID  // Assign string
Note: %d format still works with string (prints as-is)
Better to use %s for strings:
g.sendAnnouncment(fmt.Sprintf("GameMaster left, assigned <@%s> as new game master.", v.ID), false)
 true
}

func (g *Game) findPlayer(id string) *Player {  // CHANGED: int64 → string
for _, v := range g.Players {
v.ID == id {  // String comparison (works the same)
 v
 nil
}
```

---

## Key Changes Summary

1. **Type Changes**: 
   - `int64` → `string` for all ID fields (5 changes in Player struct)
   - Function parameters: `id int64` → `id string` (3 functions)

2. **Comparisons**: 
   - `==` operator works identically for strings and int64
   - No logic changes needed

3. **Printf Format**:
   - `%d` → `%s` in format strings (1 change)
   - Can keep `%d` temporarily as Go will print strings anyway

4. **API Calls**:
   - `g.Session.UserChannelCreate(id)` - Works with both, just type changes
   - Return values automatically match new types

---

## What Stays the Same

- Logic flow: identical
- Conditionals: identical  
- Error handling: identical
- Function structure: identical

**Migration is mostly mechanical type changes!**

---

## Build Verification

After changes, the code must compile without errors:

```bash
# This would fail initially due to type mismatches in other files:
go build ./cmd/cardsagainstdiscord

# Typical error you'd see:
# cannot use id (variable of type string) as int64 value in argument to g.findPlayer
```

You'd need to update ALL files that call these functions to use string IDs.
