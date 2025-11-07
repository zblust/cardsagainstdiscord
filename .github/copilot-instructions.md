This repository implements a Discord Cards Against Humanity bot in Go. The notes below give actionable, project-specific guidance for an AI code assistant to be immediately productive.

1) Big-picture architecture
- Core package: top-level Go package `cardsagainstdiscord` contains game logic (`game.go`, `manager.go`, `cad.go`) and card packs (`deck_*.go`).
- Entrypoint: `cmd/cardsagainstdiscord/main.go` creates a `discordgo` session and a `GameManager`, registers handlers, and opens the gateway.
- Runtime model: `GameManager` holds active `Game` instances keyed by channel and player IDs. `Game` contains players, available cards, and per-game goroutines (ticker, reaction/message handlers).

2) How packs are added and how to inspect them
- Packs are registered via init() functions in `deck_*.go` files which call `AddPack(&CardPack{...})`. Example: `deck_bluebox.go`.
- Each `CardPack` has `Prompts` ([]*PromptCard) and `Responses` ([]ResponseCard). Prompt cards may contain multiple `%s` placeholders; `cad.go` sets `NumPick` accordingly during `AddPack`.

3) Developer workflows & run commands
- Run the bot locally by setting the Discord token in `DG_TOKEN` and running the command in `cmd/cardsagainstdiscord`:

  DG_TOKEN="Bot <token>" go run ./cmd/cardsagainstdiscord

- The README shows legacy `go get` install instructions; prefer `go run` or `go build` with current Go tooling.
- The main process also starts pprof on :7447 for profiling (see `main.go`).

4) Key patterns and conventions (project-specific)
- Session provider: `SessionProvider` / `StaticSessionProvider` in `cad.go` is used so the `GameManager` can obtain a `discordgo.Session` per guild. Tests or tools should pass a `StaticSessionProvider` with a mock session.
- Game lifecycle: `Game.Created()` starts goroutines (`runTicker`) and adds reactions to the menu message. `Game.stopch` is closed to terminate the ticker.
- Concurrency: `Game` and `GameManager` extensively use `sync.RWMutex`. When modifying maps (ActiveGames) or game fields, follow the same lock patterns. Prefer using the public helper methods on `GameManager` to mutate state (CreateGame, PlayerTryJoinGame, RemoveGame) to keep invariants.
- Reaction-driven UX: game flow is driven by Discord reactions and DM messages. Reaction handlers route through `GameManager.HandleReactionAdd` and message inputs through `HandleMessageCreate`.

5) Integration points & external deps
- Uses `github.com/jonas747/discordgo`, `github.com/jonas747/dcmd`, `github.com/jonas747/dstate`.
- The bot relies on Discord gateway events and user DMs to run games. Changes to message or reaction handling must update both `cmd/cardsagainstdiscord/main.go` wiring and `GameManager` handlers.

6) Useful files to inspect when changing behavior
- `cmd/cardsagainstdiscord/main.go` — startup, command wiring, and handler registration.
- `manager.go` — game registry, creation/removal, reaction/message routing.
- `game.go` — game loop, tick logic, player management, and message sending.
- `cad.go` — card types, pack registration helpers, error types.
- `deck_*.go` — examples of pack definitions and how placeholders are used.

7) Small examples to copy-paste
- Add a new pack: create `deck_myexpansion.go` with an `init()` that constructs a `CardPack` and calls `AddPack(pack)`. See `deck_bluebox.go` for format.
- Create a test GameManager with a fake session: instantiate `NewGameManager(&StaticSessionProvider{Session: session})` and call `CreateGame(...)`.

8) What to avoid
- Don't modify `ActiveGames` map without holding `GameManager`'s lock or without using the provided public manager helpers; doing so breaks lookup invariants used by reaction routing.
- Avoid blocking calls on the main goroutine; the game logic expects reaction/message handlers to return quickly (they spawn goroutines where needed).

If any part of the runtime wiring, build instructions, or pack registration is unclear, tell me which area you'd like expanded and I will iterate.
