package messages

func EncodePitchwheel(msg map[string]interface{}) (b []byte) {
	b = make([]byte, 3)
	pitch := msg["pitch"].(int) - MIN_PITCHWHEEL
	b[0] = 0xe0 | msg["channel"].(byte)
	b[1] = byte(pitch & 0x7f)
	b[2] = byte(pitch >> 7)
	return
}

func EncodeSysex(msg map[string]interface{}) (b []byte) {
	data := msg["data"].([]byte)
	b = make([]byte, len(data)+2)
	b[0] = 0xf0
	copy(b[1:len(b)-1], data)
	b[len(b)-1] = 0xf7
	return
}

func EncodeQuarterFrame(msg map[string]interface{}) (b []byte) {
	b = make([]byte, 2)
	b[0] = 0xf1
	b[1] = msg["frame_type"].(byte)<<4 | msg["frame_value"].(byte)
	return b
}

func EncodeSongpos(data map[string]interface{}) (b []byte) {
	b = make([]byte, 3)
	pos := data["pos"].(int)
	b[0] = 0xf2
	b[1] = byte(pos & 0x7f)
	b[2] = byte(pos >> 7)
	return
}

func EncodeNoteOff(msg map[string]interface{}) (b []byte) {
	b = make([]byte, 3)
	b[0] = 0x80 | msg["channel"].(byte)
	b[1] = msg["note"].(byte)
	b[2] = msg["velocity"].(byte)
	return
}

func EncodeNoteOn(msg map[string]interface{}) (b []byte) {
	b = make([]byte, 3)
	b[0] = 0x90 | msg["channel"].(byte)
	b[1] = msg["note"].(byte)
	b[2] = msg["velocity"].(byte)
	return
}

func EncodeControlChange(msg map[string]interface{}) (b []byte) {
	b = make([]byte, 3)
	b[0] = 0xb0 | msg["channel"].(byte)
	b[1] = msg["control"].(byte)
	b[2] = msg["value"].(byte)
	return
}

var (
	ENCODE_SPECIAL_CASES = map[string]func(map[string]interface{}) []byte{
		"pitchwheel":    EncodePitchwheel,
		"sysex":         EncodeSysex,
		"quarter_frame": EncodeQuarterFrame,
		"songpos":       EncodeSongpos,

		// These are so common that they get special cases to speed things up.
		"note_off":       EncodeNoteOff,
		"note_on":        EncodeNoteOn,
		"control_change": EncodeControlChange,
	}
)

func EncodeMessage(msg map[string]interface{}) (b []byte) {
	msgType := msg["type"].(string)
	encode, ok := ENCODE_SPECIAL_CASES[msgType]
	if ok {
		b = encode(msg)
		return
	}
	spec := SPEC_BY_TYPE[msgType]
	statusByte := spec["status_byte"].(byte)
	if statusByte >= CHANNEL_MESSAGES[0] && statusByte <= CHANNEL_MESSAGES[len(CHANNEL_MESSAGES)-1] {
		statusByte |= msg["channel"].(byte)
	}
	b = make([]byte, 0)
	b = append(b, statusByte)
	for _, v := range spec["value_names"].([]string) {
		if _, ok := msg[v]; v != "channel" && ok {
			b = append(b, msg[v].(byte))
		}
	}
	return
}
