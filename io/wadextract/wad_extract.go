package wadextract

import (
	"errors"
	"fmt"
	"ltk/gamehash"
	"ltk/io/wadfile"
	"ltk/logger"
	"os"
	"strings"
)

type WadExtract struct {
	wad  *wadfile.Wad
	hash *gamehash.GameHash
}

func NewExtract(wad *wadfile.Wad, hash *gamehash.GameHash) *WadExtract {
	return &WadExtract{
		wad:  wad,
		hash: hash,
	}
}

//调用os.MkdirAll递归创建文件夹
func CreateMutiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

func (w *WadExtract) ExtractAll(location string) error {
	for hash, wadEntry := range w.wad.Entries {
		asserts, err := wadfile.NewWadEntryDataHandle(wadEntry).GetDecompressedBytes()
		if err != nil {
			return err
		}
		name := w.hash.HashTable[hash]

		// bin file tree handle
		conditionBin :=
			strings.HasPrefix(name, "data/") && strings.HasSuffix(name, ".bin")
		if conditionBin {
			binName := fmt.Sprintf("%x", hash)
			if len(binName) == 15 {
				binName = "0" + binName
			}
			name = strings.ToUpper(binName) + ".bin"
		}

		if len(name) == 0 {
			return errors.New("hash table entry not found")
		}

		index := strings.LastIndex(name, "/")
		pathAll := strings.SplitN(name, "/", index)
		parentAll := pathAll[:len(pathAll)-1]
		join := strings.Join(parentAll, "/")
		dir := location + "/" + join
		err = CreateMutiDir(dir)
		if err != nil {
			logger.Error("create dir error:", err)
			return err
		}
		write := location + "/" + name
		err2 := os.WriteFile(write, asserts, os.ModePerm)
		if err2 != nil {
			logger.Error("write file error:", err2)
			return err2
		}
	}

	return nil
}
