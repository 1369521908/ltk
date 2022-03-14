package gamehash

import (
	"testing"
)

const (
	path = "../files/hash/GAME_HASHTABLE.txt"
)

func TestGameHash_LoadGameHash(t *testing.T) {
	gamehash := NewGameHash()
	err := gamehash.LoadGameHash(path)
	if err != nil {
		t.Errorf("Error loading game hash: %s", err)
	}
}

func TestNewGameHash(t *testing.T) {

}
