package messages

import (
	"fmt"
	"mido/utils"
)

func Sysex(data []byte) (b []byte) {
	b = make([]byte, len(data)+2)
	b[0] = 0xf0
	copy(b[1:len(b)-1], data)
	b[len(b)-1] = 0xf7
	return
}

func DecodeSysexData(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["data"] = data
	return m
}

func DecodeQuarterFrameData(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["frame_type"] = data[0] >> 4
	m["frame_value"] = data[0] & 15
	return
}

func DecodeSongposData(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["pos"] = int(data[0]) | int(data[1])<<7
	return
}

func DecodePitchwheelData(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["pitch"] = int(data[0]) | int(data[1])<<7 + MIN_PITCHWHEEL
	return
}

func MakeDecodeSpecialCases() (cases map[byte]func([]byte) map[string]interface{}) {
	cases = map[byte]func([]byte) map[string]interface{}{
		0xe0: DecodePitchwheelData,
		0xf0: DecodeSysexData,
		0xf1: DecodeQuarterFrameData,
		0xf2: DecodeSongposData,
	}
	for i := 0; i < 16; i++ {
		cases[0xe0|byte(i)] = DecodePitchwheelData
	}
	return cases
}

var (
	DECODE_SPECIAL_CASES = MakeDecodeSpecialCases()
)

func DecodeDataBytes(statusByte byte, data []byte, spec map[string]interface{}) (args map[string]interface{}, err error) {
	if len(data) != spec["length"].(int)-1 {
		err = fmt.Errorf("wrong number of bytes for %s message", spec["type"])
		return
	}
	names := utils.RemoveOne(spec["value_names"].([]string), "channel")
	args = make(map[string]interface{})
	for i := 0; i < len(names) && i < len(data); i++ {
		args[names[i]] = data[i]
	}
	if statusByte >= CHANNEL_MESSAGES[0] && statusByte <= CHANNEL_MESSAGES[len(CHANNEL_MESSAGES)-1] {
		args["channel"] = statusByte & 0x0f
	}
	return
}

func DecodeMessage(msgBytes []byte, time int, check bool) (msg map[string]interface{}, err error) {
	if len(msgBytes) == 0 {
		err = fmt.Errorf("message is 0 bytes long")
		return
	}
	statusByte := msgBytes[0]
	data := msgBytes[1:]
	spec, ok := SPEC_BY_STATUS[statusByte]
	if !ok {
		err = fmt.Errorf("invalid status byte 0x%02X", statusByte)
		return
	}
	msg = map[string]interface{}{
		"type": spec["type"],
		"time": time,
	}
	if statusByte == byte(SYSEX_START) {
		if len(data) < 1 {
			err = fmt.Errorf("sysex without end byte")
			return
		}
		end := data[len(data)-1]
		data = data[:len(data)-1]
		if end != byte(SYSEX_END) {
			err = fmt.Errorf("invalid sysex end byte 0x%02X", end)
			return
		}
	}
	if check {
		err = CheckData(data)
		if err != nil {
			return
		}
	}
	var m map[string]interface{}
	if _, ok := DECODE_SPECIAL_CASES[statusByte]; ok {
		if statusByte >= CHANNEL_MESSAGES[0] && statusByte <= CHANNEL_MESSAGES[len(CHANNEL_MESSAGES)-1] {
			msg["channel"] = statusByte & 0x0f
		}
		for k, v := range DECODE_SPECIAL_CASES[statusByte](data) {
			msg[k] = v
		}
	} else {
		m, err = DecodeDataBytes(statusByte, data, spec)
		if err != nil {
			return
		}
		for k, v := range m {
			msg[k] = v
		}
	}
	return
}
