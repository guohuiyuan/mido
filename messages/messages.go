package messages

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"mido/utils"
	"sort"
	"strings"
)

type Message struct {
	Type       string  `json:"type"`
	Data       []byte  `json:"data"`
	Channel    byte    `json:"channel"`
	Control    byte    `json:"control"`
	FrameType  int     `json:"frame_type"`
	FrameValue int     `json:"frame_value"`
	Note       byte    `json:"note"`
	Pitch      int     `json:"pitch"`
	Pos        int     `json:"pos"`
	Program    byte    `json:"program"`
	Song       byte    `json:"song"`
	Value      byte    `json:"value"`
	Velocity   byte    `json:"velocity"`
	Time       float64 `json:"time"`
}

func (m *Message) Bytes() []byte {
	return EncodeMessage(m.Map())
}

func (m *Message) Hex() (res string) {
	b := m.Bytes()
	for i := 0; i < len(b); i++ {
		if i == 0 {
			res += fmt.Sprintf("%X", b[i])
		} else {
			res += fmt.Sprintf(" %X", b[i])
		}
	}
	return
}

func (m *Message) Bin() (res string) {
	res = string(m.Bytes())
	return
}

func (m *Message) Map() (mm map[string]interface{}) {
	mm, _ = utils.ToMap(m, "json")
	return
}

func (m *Message) IsRealtime() bool {
	sort.Strings(REALTIME_TYPES)
	i := sort.SearchStrings(REALTIME_TYPES, m.Type)
	return i < len(REALTIME_TYPES) && REALTIME_TYPES[i] == m.Type
}

func (m *Message) IsCc(control byte) bool {
	if m.Type != "control_change" {
		return false
	}
	return m.Control == control
}

func (m *Message) GetValueNames() []string {
	return append(SPEC_BY_TYPE[m.Type]["value_names"].([]string), "time")
}

func FromBytes(b []byte) (m Message, err error) {
	var msgMap map[string]interface{}
	var j []byte
	msgMap, err = DecodeMessage(b, 0, true)
	if err != nil {
		return
	}
	j, err = json.Marshal(msgMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &m)
	return
}

func FromHex(h string) (m Message, err error) {
	h = strings.ReplaceAll(h, " ", "")
	b, _ := hex.DecodeString(h)
	m, err = FromBytes(b)
	return
}
