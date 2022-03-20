package main

import (
	"ltk/gamehash"
	"ltk/io/wadextract"
	"ltk/io/wadfile"
	"ltk/logger"
	"net/http"
	_ "net/http/pprof"
)

func init() {

}

func main() {

	go func() {
		logger.Info(http.ListenAndServe("localhost:6060", nil))
	}()

	paths := []string{
		"files/wad/Aatrox.wad.client",
		"files/wad/Aphelios.wad.client",
		"files/wad/Yone.wad.client"}
	wad, err := wadfile.Read(paths[0])
	if err != nil {
		logger.Error(err)
		return
	}

	gamehashpath := "files/hash/hashes.game.txt"
	hash := gamehash.NewGameHash(gamehashpath)
	if err != nil {
		logger.Error("", err)
	}

	wadExtract := wadextract.NewExtract(wad, hash)

	err = wadExtract.ExtractAll("E:/test/Aatrox")
	if err != nil {
		logger.Error("", err)
	}
	select {}
}
