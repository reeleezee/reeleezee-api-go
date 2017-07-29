package misc

import "testing"

func TestMaxString(t *testing.T) {
	s := MaxString("This is a long test string", 25)
	if len(s) != 25 {
		t.Errorf("MaxString failed")
	}
	s = MaxString(nil, 25)
	if len(s) != 0 {
		t.Errorf("MaxString failed")
	}
}
func TestPseudoUuidV4(t *testing.T) {
	s := PseudoUuidV4()
	if len(s) != 36 && s[14:15] != "4" {
		t.Errorf("Invalid uuid")
	}
}
