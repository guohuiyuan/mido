package messages

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeMessage(t *testing.T) {
	// PS: 对map直接进行遍历会有切片超界的问题，不知道为什么，所以改成了k从小到大遍历
	dataBytes := []byte{0x01, 0x02, 0x03}
	keys := make([]int, 0)
	for k, _ := range SPEC_BY_STATUS {
		keys = append(keys, int(k.(byte)))
	}
	sort.Ints(keys)
	for _, v := range keys {
		msgBytes := make([]byte, 0)
		if byte(v) == 0xf0 {
			msgBytes = append(msgBytes, 0xf0)
			msgBytes = append(msgBytes, dataBytes...)
			msgBytes = append(msgBytes, 0xf7)
		} else {
			msgBytes = append(msgBytes, byte(v))
			msgBytes = append(msgBytes, dataBytes[:SPEC_BY_STATUS[byte(v)]["length"].(int)-1]...)
		}
		d, _ := DecodeMessage(msgBytes, 0, true)
		assert.Equal(t, msgBytes, EncodeMessage(d), "encode error or decode error")
	}
}
