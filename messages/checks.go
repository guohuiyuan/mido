package messages

import (
	"errors"
	"fmt"
	"sort"
)

func CheckType(msgType interface{}) (err error) {
	_, ok := SPEC_BY_TYPE[msgType]
	if !ok {
		err = errors.New(fmt.Sprintf("invalid message type %v", msgType))
	}
	return
}

func CheckChannel(channel interface{}) (err error) {
	c, ok := channel.(byte)
	if !ok {
		err = errors.New("channel must be byte")
	} else if c < 0 || c > 15 {
		err = errors.New("channel must be in range 0..15")
	}
	return
}

func CheckPos(pos interface{}) (err error) {
	p, ok := pos.(int)
	if !ok {
		err = errors.New("song pos must be int")
	} else if p < MIN_SONGPOS || p > MAX_SONGPOS {
		err = errors.New(fmt.Sprintf("song pos must be in range %d..%d", MIN_SONGPOS, MAX_SONGPOS))
	}
	return
}

func CheckPitch(pitch interface{}) (err error) {
	p, ok := pitch.(int)
	if !ok {
		err = errors.New("pichwheel value must be int")
	} else if p < MIN_PITCHWHEEL || p > MAX_PITCHWHEEL {
		err = errors.New(fmt.Sprintf("song pos must be in range %d..%d", MIN_PITCHWHEEL, MAX_PITCHWHEEL))
	}
	return
}

func CheckData(dataBytes interface{}) (err error) {
	d, ok := dataBytes.([]byte)
	if !ok {
		err = errors.New("data must be byte array")
		return
	}
	for _, v := range d {
		err = CheckDataByte(v)
		if err != nil {
			return
		}
	}
	return
}

func CheckFrameType(value interface{}) (err error) {
	v, ok := value.(int)
	if !ok {
		err = errors.New("frame_type must be int")
	} else if v < 0 || v > 7 {
		err = errors.New("frame_type must be int 0..7")
	}
	return
}

func CheckFrameValue(value interface{}) (err error) {
	v, ok := value.(int)
	if !ok {
		err = errors.New("frame_value must be int")
	} else if v < 0 || v > 15 {
		err = errors.New("frame_value must be int 0..15")
	}
	return
}

func CheckDataByte(value interface{}) (err error) {
	v, ok := value.(byte)
	if !ok {
		err = errors.New("data byte must be byte")
	} else if v < 0 || v > 127 {
		err = errors.New("data byte must be in range 0..127")
	}
	return
}

func CheckTime(time interface{}) (err error) {
	_, ok1 := time.(int)
	_, ok2 := time.(float64)
	if !ok1 && !ok2 {
		err = errors.New("time must be int or float")
	}
	return
}

var (
	CHECKS = map[string]func(interface{}) error{
		"type":        CheckType,
		"data":        CheckData,
		"channel":     CheckChannel,
		"control":     CheckDataByte,
		"frame_type":  CheckFrameType,
		"frame_value": CheckFrameValue,
		"note":        CheckDataByte,
		"pitch":       CheckPitch,
		"pos":         CheckPos,
		"program":     CheckDataByte,
		"song":        CheckDataByte,
		"value":       CheckDataByte,
		"velocity":    CheckDataByte,
		"time":        CheckTime,
	}
)

func CheckValue(name string, value interface{}) (err error) {
	err = CHECKS[name](value)
	return
}

func CheckMsg(msg map[string]interface{}) (err error) {
	spec, ok := SPEC_BY_TYPE[msg["type"]]
	if !ok {
		err = errors.New(fmt.Sprintf("unknown message type %v", msg["type"]))
		return
	}
	for k, v := range msg {
		a := spec["attribute_names"].([]string)
		sort.Strings(a)
		i := sort.SearchStrings(a, k)
		if i >= len(a) || a[i] != k {
			err = errors.New(fmt.Sprintf("%v message has no attribute %v", spec["type"], k))
			return
		}
		err = CheckValue(k, v)
		if err != nil {
			return
		}
	}
	return
}
