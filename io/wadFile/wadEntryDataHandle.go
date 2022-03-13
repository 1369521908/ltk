package wadFile

import (
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/klauspost/compress/zstd"
	"io"
)

type WadEntryDataHandle struct {
	wadEntry *WadEntry
}

func NewWadEntryDataHandle(wadEntry *WadEntry) *WadEntryDataHandle {
	return &WadEntryDataHandle{wadEntry}
}

func (h *WadEntryDataHandle) GetCompressedBytes() ([]byte, error) {

	file := h.wadEntry.wad.File
	entry := h.wadEntry

	// reset
	reset, err := file.Seek(0x00, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	defer file.Seek(reset, io.SeekStart)

	// Seek to entry data
	_, err = file.Seek(int64(entry._dataOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	// Read compressed data to bytes
	compressedData := make([]byte, entry.CompressedSize)
	if _, err := file.Read(compressedData); err != nil {
		return nil, err
	}

	switch entry.Type {
	case GZip:
		fallthrough
	case ZStandard:
		fallthrough
	case Uncompressed:
		return compressedData, nil
	case FileRedirection:
		return nil, errors.New("Cannot open a handle to a File Redirection")
	default:
		return nil, errors.New("Invalid Wad Entry type: " + string(entry.Type))
	}

}

func (h *WadEntryDataHandle) GetDecompressedBytes() ([]byte, error) {

	file := h.wadEntry.wad.File
	entry := h.wadEntry

	// reset
	reset, err := file.Seek(0x00, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	defer file.Seek(reset, io.SeekStart)

	// Seek to entry data
	_, err = file.Seek(int64(entry._dataOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	// Read compressed data to bytes
	compressedData := make([]byte, entry.CompressedSize)
	if _, err := file.Read(compressedData); err != nil {
		return nil, err
	}

	switch entry.Type {
	case GZip:
		// TODO need test
		uncompressedData := make([]byte, entry.UncompressedSize)
		gzipData, err := gzip.NewReader(bytes.NewReader(compressedData))
		if err != nil {
			return nil, err
		}
		uncompressedData = gzipData.Extra
		return uncompressedData, nil
	case ZStandard:
		decoder, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))
		if err != nil {
			return nil, err
		}
		uncompressedData, err := decoder.DecodeAll(compressedData, nil)
		if err != nil {
			return nil, err
		}
		return uncompressedData, nil
	case Uncompressed:
		return compressedData, nil
	case FileRedirection:
		return nil, errors.New("Cannot open a handle to a File Redirection")
	default:
		return nil, errors.New("Invalid Wad Entry type: " + string(entry.Type))
	}
}
