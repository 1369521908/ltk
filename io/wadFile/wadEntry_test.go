package wadFile

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	a := int64(0xffff16)
	b := int32(a)

	fmt.Println(b)
}
