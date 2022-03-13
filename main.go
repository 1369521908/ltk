package main

import (
	"fmt"
	"log"
	"ltk/io/wadFile"
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
		"files/wadFile/Aatrox.wadFile.client",
		"files/wadFile/Aphelios.wadFile.client",
		"files/wadFile/Yone.wadFile.client"}
	wad, err := wadFile.Read(paths[0])
	if err != nil {
		return
	}

	for k, _ := range wad.Entries {
		xxhash := fmt.Sprintf("%x", k)
		logger.Info(xxhash)
	}

	select {}
}
