package util

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"

	"github.com/Chalin-Shi/gout/libs/setting"
)

func Encrypt(origin string, algorithm string) string {
	var h hash.Hash
	if algorithm == "md5" {
		h = md5.New()
	} else {
		h = sha256.New()
	}
	io.WriteString(h, origin)
	io.WriteString(h, setting.Secret)
	return fmt.Sprintf("%x", h.Sum(nil))
}
