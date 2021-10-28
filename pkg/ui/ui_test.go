package ui

import (
	"fmt"
	"game_of_life/pkg/game"
	"testing"
)

func TestNew(t *testing.T) {
	u := New(game.GameWrapper{})
	fmt.Println(u)
}
