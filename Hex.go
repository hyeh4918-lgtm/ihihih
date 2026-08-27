package main

import "strconv"

func ConvertHex(hexStr string) (string, bool) {
	val, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return hexStr, false
	}
	return strconv.FormatInt(val, 10), true
}