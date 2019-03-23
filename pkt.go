package tlv

import (
	"io"
	"log"
	"strings"
	"strconv"
	//"github.com/tarm/serial"
)

// TLV Structure
type TLV struct {
	Type int // Type is agreed on by application
	Len int // Len is in bytes, includes 2 bytes for header
	Data []byte // byte slice Len - 2

	Reader io.Reader
	Writer io.Writer
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

// ReadLoop will loop around reading packets as they arrive
// sending them to the appriopriate call back
func ReadLoop(reader io.Reader) error {

	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			continue
		}
		incoming := string(buf[:n])
		incoming = strings.Trim(incoming, "\n\r")
		parsed := strings.Split(incoming, ":")
		switch parsed[0] {
		case "cm":
			distance := parsed[1]
			i, err := strconv.Atoi(distance)
			if err != nil {
				log.Fatal(err)
			}
			//dst := int(distance);
			log.Println(i)

		default:
			log.Printf("unknown command %v", parsed[0])
		}
	}
}
