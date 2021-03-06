package wadfile

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"ltk/logger"
	"os"
	"time"
)

type Wad struct {
	HEADER_SIZE_V3 int // default 272
	signature      []byte
	Entries        map[uint64]*WadEntry
	File           *os.File
	br             *io.Reader
	_leaveOpen     bool // not used
	_isDisposed    bool
	dataChecksum   uint64
	FileCount      uint32
}

func (w *Wad) Signature() string {
	marshal, _ := json.Marshal(w.signature)
	s := string(marshal)
	return s
}

func Read(wadPath string) (*Wad, error) {
	start := time.Now()

	file, err := os.OpenFile(wadPath, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	br := bufio.NewReader(file)

	wad := &Wad{
		HEADER_SIZE_V3: 272, // fixed 272
		Entries:        map[uint64]*WadEntry{},
		File:           file,
	}

	magic := make([]byte, 2) // RW
	if _, err := br.Read(magic); err != nil {
		return nil, err
	}
	if string(magic) != "RW" {
		return nil, errors.New("InvalidFileSignatureException")
	}

	major_ := make([]byte, 1) // 3
	if _, err := br.Read(major_); err != nil {
		return nil, err
	}
	major := uint8(major_[0])

	minor_ := make([]byte, 1) // 1
	if _, err := br.Read(minor_); err != nil {
		return nil, err
	}
	minor := uint8(minor_[0])

	if major > 3 {
		return nil, errors.New("UnsupportedFileVersion")
	}

	// v2 version maybe not work
	var signature []byte
	// probably not "dataChecksum"
	var dataChecksum []byte

	if major == 2 {
		ecdsaLength_ := make([]byte, 1)
		if _, err := br.Read(ecdsaLength_); err != nil {
			return nil, err
		}
		ecdsaLength := ecdsaLength_[0]

		signature_ := make([]byte, ecdsaLength)
		if _, err := br.Read(signature_); err != nil {
			return nil, err
		}
		signature = signature_
		// unknown fixed 83 - ecdsaLength
		unknow := make([]byte, 83-ecdsaLength)
		if _, err := br.Read(unknow); err != nil {
			return nil, err
		}

		dataChecksum_ := make([]byte, 8)
		if _, err := br.Read(dataChecksum_); err != nil {
			return nil, err
		}
		dataChecksum = dataChecksum_
	}

	if major == 3 {
		signature_ := make([]byte, 256)
		if _, err := br.Read(signature_); err != nil {
			return nil, err
		}

		dataChecksum_ := make([]byte, 8)
		if _, err := br.Read(dataChecksum_); err != nil {
			return nil, err
		}
		dataChecksum = dataChecksum_
		signature = signature_
	}

	wad.signature = signature
	// not used
	wad.dataChecksum = binary.LittleEndian.Uint64(dataChecksum)

	if major == 1 || major == 2 {
		tocStartOffset := make([]byte, 2)
		if _, err := br.Read(tocStartOffset); err != nil {
			return nil, err
		}

		tocFileEntrySize := make([]byte, 2)
		if _, err := br.Read(tocFileEntrySize); err != nil {
			return nil, err
		}
	}

	// WadEntry count
	fileCount_ := make([]byte, 4)
	if _, err := br.Read(fileCount_); err != nil {
		return nil, err
	}
	fileCount := binary.LittleEndian.Uint32(fileCount_)
	wad.FileCount = fileCount

	for i := uint32(0); i < fileCount; i++ {
		entry := NewWadEntry(wad, br, major, minor)
		if _, exist := wad.Entries[entry.XXHash]; exist {
			return nil, errors.New("Tried to read a Wad Entry with the same path gamehash as an already existing entry: " +
				fmt.Sprintf("%x", entry.XXHash))
		}
		wad.Entries[entry.XXHash] = entry
	}

	logger.Info("Wad file read in " + time.Since(start).String())
	return wad, nil
}
