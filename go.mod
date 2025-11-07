module github.com/jonas747/cardsagainstdiscord

go 1.13

require (
	github.com/jonas747/dcmd v1.1.0
	github.com/jonas747/discordgo v1.1.9
	github.com/jonas747/dstate v1.0.4
	github.com/pkg/errors v0.8.1
)

// Some upstream modules still declare the old module path (github.com/jonas747/*).
// Map those canonical paths to the botlabs-gg repos which host compatible code.
replace github.com/jonas747/dcmd => github.com/botlabs-gg/dcmd v1.1.0

replace github.com/jonas747/discordgo => github.com/botlabs-gg/discordgo v1.1.9

replace github.com/jonas747/dstate => github.com/botlabs-gg/dstate v1.0.4

// Local fallback to a compatible fork (cloned into tmp/dutil). This avoids relying
// on the original github.com/jonas747/dutil repo which no longer exists.
replace github.com/jonas747/dutil => github.com/ereti/dutil v0.0.2
