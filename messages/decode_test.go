package messages

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysex(t *testing.T) {
	msg := map[string]interface{}{
		"type": "sysex",
		"data": []byte{0x01, 0x02, 0x03},
		"time": 0,
	}
	m, _ := DecodeMessage([]byte{0xf0, 0x01, 0x02, 0x03, 0xf7}, 0, true)
	assert.True(t, reflect.DeepEqual(m, msg), "sysex decode error")
}

func TestChannel(t *testing.T) {
	m, _ := DecodeMessage([]byte{0x91, 0x00, 0x00}, 0, true)
	assert.Equal(t, uint8(1), m["channel"], "channel decode error")
}

func TestSysexEnd(t *testing.T) {
	_, err := DecodeMessage([]byte{0xf0, 0x00, 0x12}, 0, true)
	assert.EqualError(t, err, "invalid sysex end byte 0x12", "sysex end error")
}

func TestZeroByte(t *testing.T) {
	_, err := DecodeMessage([]byte{}, 0, true)
	assert.EqualError(t, err, "message is 0 bytes long", "zero byte error")
}

func TestTooFewBytes(t *testing.T) {
	_, err := DecodeMessage([]byte{0x90}, 0, true)
	assert.EqualError(t, err, "wrong number of bytes for note_on message", "too few bytes error")
}

func TestTooManyBytes(t *testing.T) {
	_, err := DecodeMessage([]byte{0x90, 0x00, 0x00, 0x00}, 0, true)
	assert.EqualError(t, err, "wrong number of bytes for note_on message", "too many bytes error")
}

func TestInvalidStatus(t *testing.T) {
	_, err := DecodeMessage([]byte{0x00}, 0, true)
	assert.EqualError(t, err, "invalid status byte 0x00", "invalid status error")
}

func TestSysexWithoutStopByte(t *testing.T) {
	_, err := DecodeMessage([]byte{0xf0}, 0, true)
	assert.EqualError(t, err, "sysex without end byte", "sysex without stop byte error")
	_, err = DecodeMessage([]byte{0xf0, 0x00}, 0, true)
	assert.EqualError(t, err, "invalid sysex end byte 0x00", "sysex without stop byte error")
}
