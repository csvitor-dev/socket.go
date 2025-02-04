package socket

import "strings"

type DefaultHandler struct {
}

func (h *DefaultHandler) Push(request []byte) []byte {
	result := string(request)

	result = strings.ToUpper(result)

	return []byte(result)
}