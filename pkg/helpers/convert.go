package helpers

import (
	"fmt"
	"strconv"
)

func String2Uint(str *string) *uint {
	if str == nil {
		return nil
	}

	u64, err := strconv.ParseUint(*str, 10, 64)
	if err != nil {
		return nil
	}

	result := uint(u64)
	return &result
}

func Uint2String(uintVal *uint) *string {
	if uintVal == nil {
		return nil
	}
	result := fmt.Sprint(*uintVal)
	return &result
}
