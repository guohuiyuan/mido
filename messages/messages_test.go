package messages

import "testing"

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
	// pitchwheel
	b = []byte{0xE0, 00, 00}
	m, err = FromBytes(b)
	t.Logf("m1:%+v,err:%v", m, err)
}

func TestFromHex(t *testing.T) {
	var b []byte
	var m Message
	var err error
	// pitchwheel
	m, err := FromHex("E0 00 00")
	t.Logf("m:%+v,err:%v", m, err)
}
