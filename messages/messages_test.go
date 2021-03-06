package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	m := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m.Bytes():%v", m.Bytes())
	t.Logf("m:%+v", m)
}

func TestHex(t *testing.T) {
	m := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m.Hex():%v", m.Hex())
}

func TestBin(t *testing.T) {
	m := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m.Bin():%v", m.Bin())
}

func TestMap(t *testing.T) {
	m := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m.Map():%v", m.Map())
}

func TestIsRealtime(t *testing.T) {
	m1 := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m1.IsRealtime():%v", m1.IsRealtime())
	m2 := Message{
		Type:     "tune_request",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m2.IsRealtime():%v", m2.IsRealtime())
}

func TestIsCc(t *testing.T) {
	m1 := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m1.IsCc(0):%v", m1.IsCc(0))
	m2 := Message{
		Type:     "control_change",
		Note:     60,
		Velocity: 64,
		Control:  1,
	}
	t.Logf("m2.IsCc(0):%v", m2.IsCc(0))
	t.Logf("m2.IsCc(1):%v", m2.IsCc(1))
}

func TestGetValueNames(t *testing.T) {
	m := Message{
		Type:     "note_on",
		Note:     60,
		Velocity: 64,
	}
	t.Logf("m.GetValueNames():%v", m.GetValueNames())
}

func TestFromBytes(t *testing.T) {
	var b []byte
	var m Message
	var err error
	b = []byte{0xE0, 00, 00}
	m, err = FromBytes(b)
	t.Logf("m1:%+v,err:%v", m, err)
}

func TestFromHex(t *testing.T) {
	var m Message
	var err error
	m, err = FromHex("E0 7F 7F")
	t.Logf("m:%+v,err:%v", m, err)
}

func TestEncodePitchwheel(t *testing.T) {
	m1 := Message{
		Type:  "pitchwheel",
		Pitch: MIN_PITCHWHEEL,
	}
	assert.Equal(t, "E0 00 00", m1.Hex(), "pitchwheel encode error")
	m2 := Message{
		Type:  "pitchwheel",
		Pitch: 0,
	}
	assert.Equal(t, "E0 00 40", m2.Hex(), "pitchwheel encode error")
	m3 := Message{
		Type:  "pitchwheel",
		Pitch: MAX_PITCHWHEEL,
	}
	assert.Equal(t, "E0 7F 7F", m3.Hex(), "pitchwheel encode error")
}

func TestDecodePitchwheel(t *testing.T) {
	var m Message
	m, _ = FromHex("E0 00 00")
	assert.Equal(t, MIN_PITCHWHEEL, m.Pitch, "pitchwheel decode error")
	m, _ = FromHex("E0 00 40")
	assert.Equal(t, 0, m.Pitch, "pitchwheel decode error")
	m, _ = FromHex("E0 7F 7F")
	assert.Equal(t, MAX_PITCHWHEEL, m.Pitch, "pitchwheel decode error")
}

func TestEncodeSongpos(t *testing.T) {
	m := Message{Type: "songpos", Pos: MIN_SONGPOS}
	assert.Equal(t, "F2 00 00", m.Hex(), "songpos encode error")
	m.Pos = MAX_SONGPOS
	assert.Equal(t, "F2 7F 7F", m.Hex(), "songpos encode error")
}

func TestDecodeSongpos(t *testing.T) {
	var m Message
	m, _ = FromHex("F2 00 00")
	assert.Equal(t, MIN_SONGPOS, m.Pos, "songpos decode error")
	m, _ = FromHex("F2 7F 7F")
	assert.Equal(t, MAX_SONGPOS, m.Pos, "songpos decode error")
}
