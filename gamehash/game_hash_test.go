package gamehash

import (
	"testing"
)

const (
	path = "../files/hash/hashes.game.txt"
)

func TestGameHash_LoadGameHash(t *testing.T) {
	gamehash := NewGameHash(path)
	t.Log(len(gamehash.HashTable))
}

func TestNewGameHash(t *testing.T) {

}
