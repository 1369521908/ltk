package wadextract

import (
	"ltk/gamehash"
	"ltk/io/wadfile"
	"testing"
)

var wadpath = "../../files/wad/Aatrox.wad.client"
var gamehashpath = "../../files/hash/hashes.game.txt"

func TestNewExtract(t *testing.T) {

}

func TestWadExtract_ExtractAll(t *testing.T) {
	wad, err := wadfile.Read(wadpath)

	if err != nil {
		t.Error(err)
	}

	hash := gamehash.NewGameHash(gamehashpath)
	if err != nil {
		t.Error(err)
	}

	wadExtract := NewExtract(wad, hash)

	err = wadExtract.ExtractAll("E:/test/Aatrox")
	if err != nil {
		t.Error(err)
	}

}
