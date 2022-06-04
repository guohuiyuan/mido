package messages

import (
	"testing"
)

func TestCheckTime(t *testing.T) {
	t.Logf("CheckTime(1):%s", CheckTime(1))
	t.Logf("CheckTime(1.5):%s", CheckTime(1.5))
	t.Logf("CheckTime(0xff):%s", CheckTime(0xff))
	t.Logf("CheckTime(nil):%s", CheckTime(nil))
	t.Logf("CheckTime(\"abc\"):%s", CheckTime("abc"))
}
