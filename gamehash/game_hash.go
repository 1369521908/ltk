package gamehash

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"github.com/cespare/xxhash"
	"hash"
	"io"
	"os"
	"strings"
)

type Hash struct {
	hash.Hash
}

type GameHash struct {
	HashTable map[uint64]string
	hash      Hash
}

func NewGameHash(path string) *GameHash {
	gameHash := &GameHash{
		HashTable: make(map[uint64]string),
		hash:      Hash{xxhash.New()},
	}
	gameHash.loadGameHash(path)

	return gameHash
}

func (h *GameHash) loadGameHash(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		split := strings.Split(string(line), " ")

		var xxHash uint64
		var name string
		if len(split) == 1 {
			sum := h.hash.Sum([]byte(split[0]))
			xxHash = binary.LittleEndian.Uint64(sum)
		} else {
			for i := 1; i < len(split); i++ {
				name += split[i]
				if i+1 != len(split) {
					name += " "
				}
			}
			decodeString, err := hex.DecodeString(split[0])
			if err != nil {
				return err
			}
			xxHash = binary.BigEndian.Uint64(decodeString)

		}

		h.HashTable[xxHash] = name
	}

	return nil
}
