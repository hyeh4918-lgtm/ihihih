package main

import "strconv"

func ConvertBin(binStr string) (string, bool) {
	val, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr, false
	}
	return strconv.FormatInt(val, 10), true
}