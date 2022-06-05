package messages

import "testing"

func TestSysex(t *testing.T) {
	data := []byte{0xf0, 0x01, 0x02, 0x03, 0xf7}
	msg := map[string]interface{}{
		"type": "sysex",
		"data": []byte{0x01, 0x02, 0x03},
		"time": 0,
	}
	t.Logf("data:%v", data)
	t.Logf("msg:%v", msg)
}
