package wadfile

import (
	"bufio"
	"encoding/binary"
	"ltk/logger"
)

type WadEntry struct {
	TOC_SIZE_V3         int // default 32
	XXHash              uint64
	CompressedSize      uint32
	UncompressedSize    uint32
	Type                WadEntryType
	_subChunkCount      uint32
	ChecksumType        WadEntryChecksumType
	Checksum            []byte
	FileRedirection     string
	_dataOffset         uint32
	_isDuplicated       bool
	_firstSubChunkIndex uint16
	wad                 *Wad // parent
}

type WadEntryType byte

const (
	Uncompressed        WadEntryType = 0
	GZip                WadEntryType = 1
	FileRedirection     WadEntryType = 2
	ZStandard           WadEntryType = 3
	ZStandardWithChunks WadEntryType = 4
)

type WadEntryChecksumType byte

const (
	SHA256  WadEntryChecksumType = 0
	XXHash3 WadEntryChecksumType = 1
)

func NewWadEntry(wad *Wad, br *bufio.Reader, major byte, minor byte) *WadEntry {

	w := &WadEntry{}
	w.wad = wad

	xxHash_ := make([]byte, 8)
	if _, err := br.Read(xxHash_); err != nil {
		return nil
	}
	w.XXHash = binary.LittleEndian.Uint64(xxHash_)

	_dataOffset_ := make([]byte, 4)
	if _, err := br.Read(_dataOffset_); err != nil {
		return nil
	}
	w._dataOffset = binary.LittleEndian.Uint32(_dataOffset_)

	compressedSize_ := make([]byte, 4)
	if _, err := br.Read(compressedSize_); err != nil {
		return nil
	}
	w.CompressedSize = binary.LittleEndian.Uint32(compressedSize_)

	uncompressedSize_ := make([]byte, 4)
	if _, err := br.Read(uncompressedSize_); err != nil {
		return nil
	}
	w.UncompressedSize = binary.LittleEndian.Uint32(uncompressedSize_)

	type_ := make([]byte, 1)
	if _, err := br.Read(type_); err != nil {
		return nil
	}
	entryType := WadEntryType(uint8(type_[0]))
	// ???
	w._subChunkCount = uint32(w.Type >> 4)
	// ???
	w.Type = entryType & 0xF
	_isDuplicated_ := make([]byte, 1)
	if _, err := br.Read(_isDuplicated_); err != nil {
		return nil
	}
	w._isDuplicated = uint8(_isDuplicated_[0]) > 0

	firstSubChunkIndex_ := make([]byte, 2)
	if _, err := br.Read(firstSubChunkIndex_); err != nil {
		return nil
	}
	w._firstSubChunkIndex = binary.LittleEndian.Uint16(firstSubChunkIndex_)

	if major >= 2 {
		checksum_ := make([]byte, 8)
		if _, err := br.Read(checksum_); err != nil {
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
		logger.Error("don`t support this file type yet")
	}

	return w
}
