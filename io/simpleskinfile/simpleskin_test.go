package simpleskinfile

import (
	"os"
	"testing"
)

const (
	sknfile = "../../files/skn/aatrox.skn"
)

func TestNewSimpleSkin(t *testing.T) {

	file, err := os.ReadFile(sknfile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}
	NewSimpleSkin(file, false)
}
