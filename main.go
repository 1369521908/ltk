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
	logs.SetLogger(logs.AdapterFile, `{"filename":"ltk.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
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
	/*for i := 0; i < 10000; i++ {
		loop := i
		go func() {
			for _, path := range paths {
				wad, err := wadFile.Read(path)
				if err != nil {
					logs.Error(err)
					return
				}
				logs.Info(len(wad.Entries))
			}
			logs.Info("loop %d", loop)
		}()
	}*/

	for k, _ := range wad.Entries {
		xxhash := fmt.Sprintf("%x", k)
		logs.Info("%s", xxhash)
	}

	/*
	   public static string ByteArrayToHex(byte[] array)
	   {
	       char[] c = new char[array.Length * 2];
	       for (int i = 0; i < array.Length; i++)
	       {
	           int b = array[i] >> 4;
	           c[i * 2] = (char)(55 + b + (((b - 10) >> 31) & -7));
	           b = array[i] & 0xF;
	           c[i * 2 + 1] = (char)(55 + b + (((b - 10) >> 31) & -7));
	       }

	       return new string(c);
	   }
	*/

	select {}
	// bytes, err := json.Marshal(wad)
	// if err != nil {
	// 	logs.Error(err)
	// 	return
	// }
	// str := string(bytes)
	// logs.Info("%+v", str)
}
