package hash

import (
	"github.com/speps/go-hashids"
)

func Encrypt(salt string, length int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = length
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.Encode([]int{0706})
		if err == nil {
			return e
		}
	}
	return ""
}
