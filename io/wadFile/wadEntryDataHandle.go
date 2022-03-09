package wadFile

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
)

type WadEntryDataHandle struct {
	w *WadEntry
}

func NewWadEntryDataHandle(w *WadEntry) *WadEntryDataHandle {
	return &WadEntryDataHandle{w: w}
}

func (h *WadEntryDataHandle) GetDecompressedBytes() ([]byte, error) {

	file := h.w.wad.File
	entry := h.w

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
	data := make([]byte, entry.CompressedSize)
	if _, err := file.Read(data); err != nil {
		return nil, err
	}

	switch entry.Type {
	case GZip:
		// TODO need test
		uncompressedData := make([]byte, entry.UncompressedSize)
		compressedData := make([]byte, entry.CompressedSize)

		gzipData, err := gzip.NewReader(bytes.NewReader(compressedData))
		if err != nil {
			return nil, err
		}
		uncompressedData = gzipData.Extra
		return uncompressedData, nil
	case ZStandard:
		// TODO need test

		// byte[] decompressedData = Zstd.Decompress(compressedData, this._entry.UncompressedSize);

		// return new MemoryStream(decompressedData);
		return nil, err
	case Uncompressed:
		return data, nil
		// return new MemoryStream(compressedData);
	case FileRedirection:
		return nil, errors.New("Cannot open a handle to a File Redirection")
	default:
		return nil, errors.New("Invalid Wad Entry type: " + string(entry.Type))
	}
}

func (h *WadEntryDataHandle) GetDecompressedStream() ([]byte, error) {

	file := h.w.wad.File
	entry := h.w

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
	data := make([]byte, entry.CompressedSize)
	if _, err := file.Read(data); err != nil {
		return nil, err
	}

	switch entry.Type {
	case GZip:
		fallthrough
	case ZStandard:
		fallthrough
	case Uncompressed:
		return data, nil
		// return new MemoryStream(compressedData);
	case FileRedirection:
		return nil, errors.New("Cannot open a handle to a File Redirection")
	default:
		return nil, errors.New("Invalid Wad Entry type: " + string(entry.Type))
	}

}
