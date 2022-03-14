package wadfile

type WadEntryBuilder struct {
	EntryType            WadEntryType
	PathXXHash           uint64
	data                 []byte
	CompressedSize       uint
	UncompressedSize     uint
	_isGenericDataStream bool
	_dataOffset          uint
	FileRedirection      string
	ChecksumType         WadEntryChecksumType
	Checksum             []byte
}

func NewWadEntryBuilder(wadEntry *WadEntry) *WadEntryBuilder {

	builder := &WadEntryBuilder{
		ChecksumType: wadEntry.ChecksumType,
	}
	// WithPathXXHash(entry.XXHash);
	switch wadEntry.Type {
	case Uncompressed:
	case GZip:
	case ZStandard:
	case FileRedirection:
	}
	return builder
}

func NewWadEntryBuilderByType(checksumType WadEntryChecksumType) *WadEntryBuilder {
	return &WadEntryBuilder{
		ChecksumType: checksumType,
	}
}

func (w *WadEntryBuilder) WithUncompressedData([]byte) *WadEntryBuilder {

	builder := &WadEntryBuilder{
		// ChecksumType: Uncompressed,
	}
	return builder
}
