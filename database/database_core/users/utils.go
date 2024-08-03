package users

import (
	"crypto/sha256"
	"encoding/hex"
	// database "github.com/creatorkostas/KeyDB/database/database_core"
)

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
