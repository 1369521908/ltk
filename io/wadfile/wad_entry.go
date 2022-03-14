package wadfile

import (
	"encoding/binary"
	"io"
	"os"
)

type WadEntry struct {
	TOC_SIZE_V3      int // default 32
	XXHash           uint64
	CompressedSize   uint32
	UncompressedSize uint32
	Type             WadEntryType
	ChecksumType     WadEntryChecksumType
	Checksum         []byte
	FileRedirection  string
	_dataOffset      uint32
	_isDuplicated    bool
	wad              *Wad // parent
}

type WadEntryType byte

const (
	Uncompressed    WadEntryType = 0
	GZip            WadEntryType = 1
	FileRedirection WadEntryType = 2
	ZStandard       WadEntryType = 3
	// ZStandardWithsubchunks WadEntryType = 4
)

type WadEntryChecksumType byte

const (
	SHA256  WadEntryChecksumType = 0
	XXHash3 WadEntryChecksumType = 1
)

func NewWadEntry(wad *Wad, file *os.File, major byte, minor byte) *WadEntry {

	w := &WadEntry{}
	w.wad = wad

	xxHash_ := make([]byte, 8)
	if _, err := file.Read(xxHash_); err != nil {
		return nil
	}
	w.XXHash = binary.LittleEndian.Uint64(xxHash_)

	_dataOffset_ := make([]byte, 4)
	if _, err := file.Read(_dataOffset_); err != nil {
		return nil
	}
	w._dataOffset = binary.LittleEndian.Uint32(_dataOffset_)

	compressedSize_ := make([]byte, 4)
	if _, err := file.Read(compressedSize_); err != nil {
		return nil
	}
	w.CompressedSize = binary.LittleEndian.Uint32(compressedSize_)

	uncompressedSize_ := make([]byte, 4)
	if _, err := file.Read(uncompressedSize_); err != nil {
		return nil
	}
	w.UncompressedSize = binary.LittleEndian.Uint32(uncompressedSize_)

	type_ := make([]byte, 1)
	if _, err := file.Read(type_); err != nil {
		return nil
	}
	w.Type = WadEntryType(uint8(type_[0]))

	_isDuplicated_ := make([]byte, 1)
	if _, err := file.Read(_isDuplicated_); err != nil {
		return nil
	}
	w._isDuplicated = uint8(_isDuplicated_[0]) > 0

	// pad
	pad_ := make([]byte, 2)
	if _, err := file.Read(pad_); err != nil {
		return nil
	}

	if major > 2 {
		checksum_ := make([]byte, 8)
		if _, err := file.Read(checksum_); err != nil {
			return nil
		}
		w.Checksum = checksum_

		if major == 3 && minor == 1 {
			w.ChecksumType = XXHash3
		} else {
			w.ChecksumType = SHA256
		}
	}

	if w.Type == FileRedirection {
		// TODO need test
		cur, err := file.Seek(0x00, io.SeekCurrent)
		if err != nil {
			return nil
		}

		_, err = file.Seek(int64(w._dataOffset), io.SeekStart)
		if err != nil {
			return nil
		}

		fileRedirection := make([]byte, 4)
		if _, err := file.Read(fileRedirection); err != nil {
			return nil
		}
		w.FileRedirection = string(fileRedirection)

		_, err = file.Seek(cur, io.SeekStart)
		if err != nil {
			return nil
		}
	}

	return w
}
