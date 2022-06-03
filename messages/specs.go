package messages

import (
	"errors"
	"math"
	"mido/utils"
	"sort"
)

var (
	CHANNEL_MESSAGES  = []byte{0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f, 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf, 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf, 0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef}
	COMMON_MESSAGES   = []byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7}
	REALTIME_MESSAGES = []byte{0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff}

	SYSEX_START = 0xf0
	SYSEX_END   = 0xf7

	// Pitchwheel is a 14 bit signed integer
	MIN_PITCHWHEEL = -8192
	MAX_PITCHWHEEL = 8191

	// Song pos is a 14 bit unsigned integer
	MIN_SONGPOS = 0
	MAX_SONGPOS = 16383
)

// type Message struct {
// 	StatusByte     byte     `json:"status_byte"`
// 	Type           string   `json:"type"`
// 	ValueNames     []string `json:"value_names"`
// 	AttributeNames []string `json:"attribute_names"`
// 	Length         int      `json:"length"`
// 	Time           int      `json:"time"`
// }

func DefMsg(statusByte byte, msgType string, valueNames []string, length int) (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["status_byte"] = statusByte
	m["type"] = msgType
	sort.Strings(valueNames)
	m["value_names"] = valueNames
	c := []string{"type", "time"}
	m["attribute_names"] = utils.RemoveDuplicate(append(valueNames, c...)...)
	m["length"] = length
	return
}

var (
	SPECS = []map[string]interface{}{
		DefMsg(0x80, "note_off", []string{"channel", "note", "velocity"}, 3),
		DefMsg(0x90, "note_on", []string{"channel", "note", "velocity"}, 3),
		DefMsg(0xa0, "polytouch", []string{"channel", "note", "value"}, 3),
		DefMsg(0xb0, "control_change", []string{"channel", "control", "value"}, 3),
		DefMsg(0xc0, "program_change", []string{"channel", "program"}, 2),
		DefMsg(0xd0, "aftertouch", []string{"channel", "value"}, 2),
		DefMsg(0xe0, "pitchwheel", []string{"channel", "pitch"}, 3),
		// System common messages.
		// 0xf4 and 0xf5 are undefined.
		DefMsg(0xf0, "sysex", []string{"data"}, math.MaxInt64),
		DefMsg(0xf1, "quarter_frame", []string{"frame_type", "frame_value"}, 2),
		DefMsg(0xf2, "songpos", []string{"pos"}, 3),
		DefMsg(0xf3, "song_select", []string{"song"}, 2),
		DefMsg(0xf6, "tune_request", []string{}, 1),
		// System real time messages.
		// 0xf9 and 0xfd are undefined.
		DefMsg(0xf8, "clock", []string{}, 1),
		DefMsg(0xfa, "start", []string{}, 1),
		DefMsg(0xfb, "continue", []string{}, 1),
		DefMsg(0xfc, "stop", []string{}, 1),
		DefMsg(0xfe, "active_sensing", []string{}, 1),
		DefMsg(0xff, "reset", []string{}, 1),
	}
)

func MakeSpecLookups(specs []map[string]interface{}) (lookup, byStatus, byType map[interface{}]map[string]interface{}) {
	lookup = make(map[interface{}]map[string]interface{})
	byStatus = make(map[interface{}]map[string]interface{})
	byType = make(map[interface{}]map[string]interface{})
	for _, v := range specs {
		msgType := v["type"].(string)
		statusByte := v["status_byte"].(byte)
		byType[msgType] = v
		if statusByte >= CHANNEL_MESSAGES[0] && statusByte <= CHANNEL_MESSAGES[len(CHANNEL_MESSAGES)-1] {
			for i := 0; i < 16; i++ {
				byStatus[statusByte|byte(i)] = v
			}
		} else {
			byStatus[statusByte] = v
		}
	}
	for k, v := range byStatus {
		lookup[k] = v
	}
	for k, v := range byType {
		lookup[k] = v
	}
	return
}

var (
	SPEC_LOOKUP, SPEC_BY_STATUS, SPEC_BY_TYPE = MakeSpecLookups(SPECS)
	REALTIME_TYPES                            = []string{"tune_request", "clock", "start", "continue", "stop"}
	DEFAULT_VALUES                            = map[string]interface{}{
		"channel":     0,
		"control":     0,
		"data":        []interface{}{},
		"frame_type":  0,
		"frame_value": 0,
		"note":        0,
		"pitch":       0,
		"pos":         0,
		"program":     0,
		"song":        0,
		"value":       0,
		"velocity":    64,
		"time":        0,
	}
)

func MakeMsg(msgType string, overrides map[string]interface{}) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	var spec map[string]interface{}
	var ok bool
	if spec, ok = SPEC_BY_TYPE[msgType]; !ok {
		err = errors.New("Unknown message type " + msgType)
		return
	}
	m["type"] = msgType
	m["time"] = DEFAULT_VALUES["time"]
	for _, v := range spec["value_names"].([]string) {
		m[v] = DEFAULT_VALUES[v]
	}
	for k, v := range overrides {
		m[k] = v
	}
	return
}
