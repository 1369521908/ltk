package helper

import (
	"strconv"
	"strings"
)

func HashToHex16(hash uint64) string {
	hex := strconv.FormatUint(hash, 16)
	for {
		if len(hex) < 16 {
			hex = "0" + hex
		} else {
			break
		}
	}
	return strings.ToUpper(hex)
}
