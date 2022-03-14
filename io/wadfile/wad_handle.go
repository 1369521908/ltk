package wadfile

type WadHandle struct {
	wad *Wad
}

func NewWadHandle(wad *Wad) *WadHandle {
	return &WadHandle{wad}
}

func (w *WadHandle) Export() error {
	return nil
}

func (w *WadHandle) Close() error {
	err := w.wad.File.Close()
	if err != nil {
		return err
	}
	return nil
}
