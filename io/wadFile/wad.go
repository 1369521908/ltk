package wadFile

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io/fs"
	"os"
)

type Wad struct {
	HEADER_SIZE_V3 int // default 272
	Signature      []byte
	Entries        map[uint64]*WadEntry
	// data           []byte
	File *os.File
	// _leaveOpen     bool // not used
	_isDisposed  bool
	DataChecksum uint64
	FileCount    uint32
}

func Read(wadPath string) (*Wad, error) {
	file, err := os.OpenFile(wadPath, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wad := &Wad{
		HEADER_SIZE_V3: 272, // fixed 272
		Entries:        map[uint64]*WadEntry{},
		File:           file,
	}

	magic := make([]byte, 2) // RW
	if _, err := file.Read(magic); err != nil {
		return nil, err
	}
	if string(magic) != "RW" {
		return nil, errors.New("InvalidFileSignatureException")
	}

	major_ := make([]byte, 1) // 3
	if _, err := file.Read(major_); err != nil {
		return nil, err
	}
	major := uint8(major_[0])

	minor_ := make([]byte, 1) // 1
	if _, err := file.Read(minor_); err != nil {
		return nil, err
	}
	minor := uint8(minor_[0])

	if major > 3 {
		return nil, errors.New("UnsupportedFileVersion")
	}

	// v2 version maybe not work
	var signature []byte
	var dataChecksum []byte
	if major == 2 {
		ecdsaLength_ := make([]byte, 1)
		if _, err := file.Read(ecdsaLength_); err != nil {
			return nil, err
		}

		ecdsaLength := ecdsaLength_[0]
		signature_ := make([]byte, ecdsaLength)
		if _, err := file.Read(signature_); err != nil {
			return nil, err
		}
		signature = signature_
		// unknown fixed 83 - ecdsaLength
		unknow := make([]byte, 83-ecdsaLength)
		if _, err := file.Read(unknow); err != nil {
			return nil, err
		}

		dataChecksum_ := make([]byte, 8)
		if _, err := file.Read(dataChecksum_); err != nil {
			return nil, err
		}
		dataChecksum = dataChecksum_
	} else if major == 3 {
		signature_ := make([]byte, 256)
		if _, err := file.Read(signature_); err != nil {
			return nil, err
		}

		dataChecksum_ := make([]byte, 8)
		if _, err := file.Read(dataChecksum_); err != nil {
			return nil, err
		}
		dataChecksum = dataChecksum_
		signature = signature_
	}

	wad.Signature = signature
	wad.DataChecksum = binary.LittleEndian.Uint64(dataChecksum)

	if major == 1 || major == 2 {
		tocStartOffset := make([]byte, 2)
		if _, err := file.Read(tocStartOffset); err != nil {
			return nil, err
		}

		tocFileEntrySize := make([]byte, 2)
		if _, err := file.Read(tocFileEntrySize); err != nil {
			return nil, err
		}
	}

	// WadEntry count
	fileCount_ := make([]byte, 4)
	if _, err := file.Read(fileCount_); err != nil {
		return nil, err
	}
	fileCount := binary.LittleEndian.Uint32(fileCount_)
	wad.FileCount = fileCount

	for i := uint32(0); i < fileCount; i++ {
		entry := NewWadEntry(wad, file, major, minor)
		if _, exist := wad.Entries[entry.XXHash]; exist {
			return nil, errors.New("Tried to read a Wad Entry with the same path hash as an already existing entry: " +
				fmt.Sprintf("%x", entry.XXHash))
		}
		wad.Entries[entry.XXHash] = entry

		bytes, err := NewWadEntryDataHandle(entry).GetDecompressedBytes()
		if err != nil {
			return nil, err
		}
		logs.Info(`bytes len is %d`, len(bytes))
	}

	return wad, nil
}
