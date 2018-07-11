package util

import (
	"crypto/sha256"
	"fmt"
	"io"

	"upgrade/backend/libs/setting"
)

func Encrypt(origin string) string {
	h := sha256.New()
	io.WriteString(h, origin)
	io.WriteString(h, setting.Secret)
	return fmt.Sprintf("%x", h.Sum(nil))
}
