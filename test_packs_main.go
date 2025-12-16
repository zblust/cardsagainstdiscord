package main

import (
"fmt"
"strings"

cad "github.com/jonas747/cardsagainstdiscord"
)

func main() {
fmt.Println("Testing pack names...")

issues := 0
for name := range cad.Packs {
strings.Contains(name, " ") {
tf("ERROR: Pack '%s' contains spaces\n", name)
strings.HasPrefix(strings.ToLower(name), "cah:") || strings.HasPrefix(strings.ToLower(name), "cah ") {
tf("ERROR: Pack '%s' has CAH prefix\n", name)
issues == 0 {
tln("SUCCESS: All pack names are properly formatted!")
tf("Total packs: %d\n", len(cad.Packs))
} else {
tf("FAILED: Found %d issues\n", issues)
}
}
