package wadExtract

import (
	"ltk/gameHash"
	"ltk/io/wadFile"
	"testing"
)

var wadpath = "../../files/wad/Aatrox.wad.client"
var gamehashpath = "../../files/hash/GAME_HASHTABLE.txt"

func TestNewExtract(t *testing.T) {

}

func TestWadExtract_ExtractAll(t *testing.T) {
	wad, err := wadFile.Read(wadpath)

	if err != nil {
		t.Error(err)
	}

	hash := gameHash.NewGameHash(gamehashpath)
	if err != nil {
		t.Error(err)
	}

	wadExtract := NewExtract(wad, hash)

	err = wadExtract.ExtractAll("E:/test/Aatrox")
	if err != nil {
		t.Error(err)
	}

}
