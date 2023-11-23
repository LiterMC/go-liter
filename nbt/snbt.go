package nbt

// TODO
/*
import (
	"regexp"
)

var unquoteStrMatcher = regexp.MustCompile(`[a-zA-Z0-9_\-\.\+]+`)

type SNBTDeocder struct {
	r io.Reader
}

func NewSNBTDeocder(r io.Reader)(d *SNBTDeocder){
	return &SNBTDeocder{
		r: r,
	}
}

func (d *SNBTDeocder)nextByte()(b byte, err error){
	var buf [1]byte
	if _, err = d.r.Read(buf[:]); err != nil {
		return
	}
	return buf[0], nil
}

func (d *SNBTDeocder)decode()(error){

}
*/
