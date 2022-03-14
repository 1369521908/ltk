package wadextract

import (
	"fmt"
	"io"
	"io/fs"
	"ltk/gamehash"
	"ltk/helper"
	"ltk/io/wadfile"
	"os"
	"strings"
)

const (
	// mapping the data/skin_0****.bin file
	packerMapping = "OBSIDIAN_PACKED_MAPPING.txt"
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
		// fill the file name
		wadEntry.FileRedirection = name

		folders := strings.Split(name, "/")
		// data/...****.bin
		if len(folders) == 2 && strings.HasSuffix(name, ".bin") {

			conditionBin := strings.HasPrefix(name, "data/") && strings.HasSuffix(name, ".bin")
			if conditionBin {
				mapping := location + "/" + packerMapping
				var file *os.File
				if isExist(mapping) {
					file, err = os.OpenFile(mapping, os.O_APPEND, fs.ModePerm)
					if err != nil {
						return err
					}
				} else {
					file, err = os.Create(mapping)
					if err != nil {
						return err
					}
				}
				indexName := helper.HashToHex16(hash) + ".bin"
				content := indexName + " = " + name + "\r\n"
				_, err := io.WriteString(file, content)
				if err != nil {
					return err
				}
				name = indexName
			}
		}

		folder := strings.Join(folders[:len(folders)-1], "/")
		dir := location + "/" + folder
		err = CreateMutiDir(dir)

		if err != nil {
			return err
		}
		write := location + "/" + name
		err = os.WriteFile(write, asserts, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
