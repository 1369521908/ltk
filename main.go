package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"log"
	"ltk/io/wadFile"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	paths := []string{
		"files/wad/Aatrox.wad.client",
		"files/wad/Aphelios.wad.client",
		"files/wad/Yone.wad.client"}
	wad, err := wadFile.Read(paths[0])
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(len(wad.Entries))

	for k, _ := range wad.Entries {
		xxhash := fmt.Sprintf("%x", k)
		logs.Info("%s", xxhash)
	}

	select {}
}
