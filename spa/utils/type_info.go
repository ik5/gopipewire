package utils

import "C"

import "strings"

func SPATypeIsA(spaType, parent string) bool {
	return spaType != "" && parent != "" && strings.Compare(spaType, parent) == 0
}
