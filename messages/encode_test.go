package messages

import (
	"testing"
)

func TestEncodeMessage(t *testing.T) {
	dataBytes := []byte{1, 2, 3}
	msgBytes := make([]byte, 0)
	for k, v := range SPEC_BY_STATUS {
		if k == 0xf0 {
			msgBytes = append(msgBytes, 0xf0)
			msgBytes = append(msgBytes, dataBytes...)
			msgBytes = append(msgBytes, 0xf7)
		} else {
			msgBytes = append(msgBytes, k.(byte))
			msgBytes = append(msgBytes, dataBytes[:v["length"].(int)-1]...)
		}
	}
}
