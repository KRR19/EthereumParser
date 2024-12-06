package hex

import (
	"strconv"
	"strings"
)

func ToDec(hex string) (int, error) {
	hex = strings.TrimPrefix(hex, "0x")

	decimal, err := strconv.ParseInt(hex, 16, 32)

	if err != nil {
		return 0, err
	}

	return int(decimal), nil
}
