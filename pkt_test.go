package tlv

import (
	"testing"
)

func TestNewTLV(t *testing.T) {
	tlv := NewTLV(1)
	if tlv.Len != 1 {
		t.Errorf("Type 1 expected Len (1) got (%d)", tlv.Len)
	}

	tlv = NewTLV(0x81)
	if tlv.Len != 2 {
		t.Errorf("Type 0x81 expected Len (2) got (%d)", tlv.Len)
	}
}
