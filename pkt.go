package tlv

// TLV Structure
type TLV struct {
	Type int // Type is agreed on by application
	Len int // Len is in bytes, includes 2 bytes for header
	Data []byte // byte slice Len - 2
}

// NewTLV will create a new data structure set by type.
func NewTLV(t int) TLV {
	tlv := TLV{
		Type: t,
		Len: 1,
	}
	if tlv.Type > 0x80 {
		tlv.Len = 2
	}
	return tlv
}
