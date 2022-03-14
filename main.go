package main

import (
	"fmt"
	"log"
	"ltk/io/wadfile"
	"ltk/logger"
	"net/http"
	_ "net/http/pprof"
)

func init() {

}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	paths := []string{
		"files/wad/Aatrox.wad.client",
		"files/wad/Aphelios.wad.client",
		"files/wad/Yone.wad.client"}
	wad, err := wadfile.Read(paths[0])
	if err != nil {
		return
	}

	for k, _ := range wad.Entries {
		xxhash := fmt.Sprintf("%x", k)
		logger.Info(xxhash)
	}

	select {}
}
